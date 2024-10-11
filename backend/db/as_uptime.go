package db

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/consts"
	"github.com/BaiMeow/NetworkMonitor/utils"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"strconv"
	"time"
)

func BatchRecordASUp(ASNs map[uint32]int, t time.Time) error {
	if !Enabled {
		return ErrDatabaseDisabled
	}
	var points []*write.Point
	for asn, links := range ASNs {
		points = append(points, influxdb2.NewPointWithMeasurement("bgp").
			AddField("up", 1).
			AddField("links", links).
			AddTag("asn", strconv.FormatUint(uint64(asn), 10)).
			SetTime(t))
	}
	if err := bgpWrite.WritePoint(context.Background(), points...); err != nil {
		log.Printf("write record fail:%v", err)
		return ErrDatabase
	}
	return nil
}

func AllASRecordAfter(after time.Time) ([]uint32, error) {
	if !Enabled {
		return nil, ErrDatabaseDisabled
	}
	var asns []uint32
	res, err := bgpQuery.Query(context.Background(),
		fmt.Sprintf(`from(bucket: "bgp-uptime")
	|> range(start: %d)    
	|> filter(fn: (r) => r["_measurement"] == "bgp" and r["_field"] == "up")
	|> unique(column: "asn")`, after.Unix()))
	if err != nil {
		log.Printf("query fail:%v", err)
		return asns, ErrDatabase
	}

	for res.Next() {
		asn, ok := res.Record().ValueByKey("asn").(string)
		if !ok {
			log.Printf("convert fail:%v", res.Record().ValueByKey("asn"))
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

func BGPASNLast10Tickers(asn uint32, last time.Time) ([]time.Time, error) {
	if !Enabled {
		return nil, ErrDatabaseDisabled
	}
	var t []time.Time
	res, err := bgpQuery.Query(context.Background(),
		fmt.Sprintf(`from(bucket: "bgp-uptime")
  |> range(start: %d, stop: %d)
  |> filter(fn: (r) => r["_measurement"] == "bgp" and r["_field"] == "up" and r["asn"] == "%d" )`,
			utils.TickOffset(last, -10).Add(time.Second*30).Unix(), last.Add(time.Second*30).Unix(), asn))
	if err != nil {
		log.Printf("query fail:%v", err)
		return t, ErrDatabase
	}
	for res.Next() {
		t = append(t, res.Record().Time())
	}
	return t, nil
}

// BGPLinks query the number of links for given ASN, startTime and window.
// param window should be either time.Minute or time.Hour
func BGPLinks(asn uint32, startTime, stopTime time.Time, window time.Duration) ([]consts.LinkTime, error) {
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
	res, err := bgpQuery.Query(context.Background(), fmt.Sprintf(`from(bucket: "bgp-uptime")
  |> range(start: %d,stop: %d)
  |> filter(fn: (r) => r["_measurement"] == "bgp" and r["_field"] == "links" and r["asn"] == "%d")
  |> aggregateWindow(every: %s, fn: max, createEmpty: true)
  |> yield(name: "max")`, startTime.Unix(), stopTime.Unix(), asn, every))
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
