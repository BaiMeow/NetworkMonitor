package db

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/consts"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"slices"
	"strconv"
	"time"
)

func BatchRecordBGP(name string, bgp *entity.BGP, t time.Time) error {
	if !Enabled {
		return ErrDatabaseDisabled
	}
	var points []*write.Point
	peerAS := make(map[uint32]map[uint32]bool)
	for _, link := range bgp.Link {
		points = append(points, influxdb.NewPointWithMeasurement(fmt.Sprintf("bgp-%s", name)).
			AddField("up", 1).
			AddTag("src", strconv.FormatUint(uint64(link.Src), 10)).
			AddTag("dst", strconv.FormatUint(uint64(link.Dst), 10)).
			SetTime(t))
		points = append(points, influxdb.NewPointWithMeasurement(fmt.Sprintf("bgp-%s", name)).
			AddField("up", 1).
			AddTag("dst", strconv.FormatUint(uint64(link.Src), 10)).
			AddTag("src", strconv.FormatUint(uint64(link.Dst), 10)).
			SetTime(t))
		if peerAS[link.Src] == nil {
			peerAS[link.Src] = make(map[uint32]bool)
		}
		peerAS[link.Src][link.Dst] = true
		if peerAS[link.Dst] == nil {
			peerAS[link.Dst] = make(map[uint32]bool)
		}
		peerAS[link.Dst][link.Src] = true
	}
	if err := networkWrite.WritePoint(context.Background(), points...); err != nil {
		log.Printf("write record fail:%v", err)
		return ErrDatabase
	}

	var peerCountPoint []*write.Point
	for asn, peers := range peerAS {
		peerCountPoint = append(peerCountPoint, influxdb.NewPointWithMeasurement(fmt.Sprintf("bgp-%s", name)).
			AddField("count", len(peers)).
			AddTag("asn", strconv.FormatUint(uint64(asn), 10)).
			SetTime(t))
	}
	if err := peerCountWrite.WritePoint(context.Background(), points...); err != nil {
		log.Printf("write peer count fail:%v", err)
		return ErrDatabase
	}

	return nil
}

func AllASRecordAfter(bgpName string, after time.Time) ([]uint32, error) {
	if !Enabled {
		return nil, ErrDatabaseDisabled
	}
	var asns []uint32
	res, err := allQuery.Query(context.Background(),
		fmt.Sprintf(`t1 =
   from(bucket: "network")
	|> range(start: %d)    
	|> filter(fn: (r) => r["_measurement"] == "%s")
  	|> group(columns: ["src"])
  	|> unique(column: "src")
  	|> keep(columns: ["src"])

t2 =
   from(bucket: "network")
	|> range(start: %d)    
	|> filter(fn: (r) => r["_measurement"] == "%s")
  	|> group(columns: ["dst"])
  	|> unique(column: "dst")
    |> map(fn: (r) => ({ r with src: r.dst }))
  	|> keep(columns: ["src"])

union(tables: [t1, t2])
|> group()
|> unique(column: "src")`, after.Unix(), bgpName, after.Unix(), bgpName))
	if err != nil {
		log.Printf("query fail:%v", err)
		return asns, ErrDatabase
	}

	for res.Next() {
		asn, ok := res.Record().ValueByKey("src").(string)
		if !ok {
			log.Printf("convert fail:%v", res.Record().ValueByKey("src"))
			return asns, ErrDatabase
		}
		asnNum, err := strconv.ParseUint(asn, 10, 32)
		if err != nil {
			log.Printf("convert fail:%v", err)
			continue
		} else {
			asns = append(asns, uint32(asnNum))
		}
	}

	return asns, nil
}

func BGPASNLast10Tickers(bgpName string, asn uint32) ([]bool, error) {
	if !Enabled {
		return nil, ErrDatabaseDisabled
	}
	res, err := allQuery.Query(context.Background(),
		fmt.Sprintf(`from(bucket: "network")
  |> range(start: -10m, stop: now())
  |> filter(fn: (r) => r["_measurement"] == "%s" and r["_field"] == "up" and r.src == "%d"))
  |> drop(columns: ["dst","src"])
  |> aggregateWindow(every: 1m, fn: max, createEmpty: true)`, bgpName, asn))
	if err != nil {
		log.Printf("query fail:%v", err)
		return nil, ErrDatabase
	}
	var t []bool
	for res.Next() {
		t = append(t, res.Record().Value() != nil)
	}
	// no data
	if len(t) == 0 {
		return slices.Repeat([]bool{false}, 10), nil
	}
	if len(t) < 10 {
		log.Printf("parse query fail:%v", err)
		return t, ErrDatabase
	}
	return t[:10], nil
}

// BGPLinks query the number of links for given ASN, startTime and window.
// param window should be either time.Minute or time.Hour
func BGPLinks(bgpName string, asn uint32, startTime, stopTime time.Time, window time.Duration) ([]consts.LinkTime, error) {
	if !Enabled {
		return nil, ErrDatabaseDisabled
	}
	every, err := parseWindow(window)
	if err != nil {
		return nil, err
	}

	res, err := allQuery.Query(context.Background(), fmt.Sprintf(`from(bucket: "network")
  |> range(start: %d, stop: %d)
  |> filter(fn: (r) => r["_measurement"] == "%s" and r.src == "%d")
  |> group(columns: ["_time"])
  |> count(column: "_value")
  |> group()
  |> aggregateWindow(every: %s, fn: max, createEmpty: true)
  |> yield(name: "max")`, startTime.Unix(), stopTime.Unix(), bgpName, asn, every))
	if err != nil {
		log.Printf("query fail:%v", err)
		return nil, ErrDatabase
	}
	return readTimeLinks(res)
}
