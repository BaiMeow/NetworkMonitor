package db

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/consts"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"strconv"
	"time"
)

func BatchRecordBGP(name string, bgp *entity.BGP, t time.Time) error {
	if !Enabled {
		return ErrDatabaseDisabled
	}
	var points []*write.Point
	for _, link := range bgp.Link {
		points = append(points, influxdb2.NewPointWithMeasurement(fmt.Sprintf("bgp-%s", name)).
			AddField("up", 1).
			AddTag("src", strconv.FormatUint(uint64(link.Src), 10)).
			AddTag("dst", strconv.FormatUint(uint64(link.Dst), 10)).
			SetTime(t))
	}
	if err := dbWrite.WritePoint(context.Background(), points...); err != nil {
		log.Printf("write record fail:%v", err)
		return ErrDatabase
	}
	return nil
}

func BatchRecordOSPF(asn uint32, ospf *entity.OSPF, t time.Time) error {
	if !Enabled {
		return ErrDatabaseDisabled
	}
	var points []*write.Point
	for _, area := range *ospf {
		for _, link := range area.Links {
			points = append(points, influxdb2.NewPointWithMeasurement(fmt.Sprintf("ospf-%d", asn)).
				AddField("up", 1).
				AddTag("src", link.Src).
				AddTag("dst", link.Dst).
				SetTime(t))
		}
	}
	if err := dbWrite.WritePoint(context.Background(), points...); err != nil {
		log.Printf("write record fail:%v", err)
		return ErrDatabase
	}
	return nil
}

func AllASRecordAfter(bgpName string, after time.Time) ([]uint32, error) {
	if !Enabled {
		return nil, ErrDatabaseDisabled
	}
	var asns []uint32
	res, err := dbQuery.Query(context.Background(),
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
	res, err := dbQuery.Query(context.Background(),
		fmt.Sprintf(`from(bucket: "network")
  |> range(start: -10m, stop: now())
  |> filter(fn: (r) => r["_measurement"] == "%s" and r["_field"] == "up" and ( r.src == "%d" or r.dst == "%d"))
  |> drop(columns: ["dst","src"])
  |> aggregateWindow(every: 1m, fn: max, createEmpty: true)`, bgpName, asn, asn))
	if err != nil {
		log.Printf("query fail:%v", err)
		return nil, ErrDatabase
	}
	var t []bool
	for res.Next() {
		t = append(t, res.Record().Value() != nil)
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

	var every string
	if window == time.Minute {
		every = "1m"
	} else if window == time.Hour {
		every = "1h"
	} else {
		log.Printf("invalid window:%v", window)
		return nil, ErrInvalidInput
	}

	var points []consts.LinkTime
	res, err := dbQuery.Query(context.Background(), fmt.Sprintf(`from(bucket: "network")
  |> range(start: %d, stop: %d)
  |> filter(fn: (r) => r["_measurement"] == "%s" and (r.src == "%d" or r.dst == "%d"))
  |> group(columns: ["_time"])
  |> count(column: "_value")
  |> group()
  |> aggregateWindow(every: %s, fn: max, createEmpty: true)
  |> yield(name: "max")`, startTime.Unix(), stopTime.Unix(), bgpName, asn, asn, every))
	if err != nil {
		log.Printf("query fail:%v", err)
		return nil, ErrDatabase
	}
	for res.Next() {
		rc := res.Record()
		var v int64
		switch value := rc.Value().(type) {
		case int64:
			v = value
		case nil:
			v = 0
		default:
			log.Printf("convert influxdb value fail:%v", rc)
		}
		points = append(points, consts.LinkTime{
			Time:  rc.Time(),
			Links: int(v),
		})
	}
	return points, nil
}
