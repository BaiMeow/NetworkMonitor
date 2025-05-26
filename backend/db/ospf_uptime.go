package db

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/consts"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"slices"
	"strconv"
	"time"
)

func OSPFLinks(ospfName string, routerId string, startTime, stopTime time.Time, window time.Duration) ([]consts.DirectedLinkTime, error) {
	if !Enabled {
		return nil, ErrDatabaseDisabled
	}
	every, err := parseWindow(window)
	if err != nil {
		return nil, err
	}

	res, err := allQuery.Query(context.Background(), fmt.Sprintf(`routerId = %s
graphName = %s

from(bucket: "network")
    |> range(start: %d, stop: %d)
    |> filter(
        fn: (r) => r["_measurement"] == graphName and (r.src == routerId or r.dst == routerId),
    )
    |> group(columns: ["_time"])
    |> reduce(
        fn: (r, accumulator) =>
            if r.src == routerId then
                {in: accumulator.in, out: accumulator.out + 1, _value: accumulator._value + 1}
            else
                {in: accumulator.in + 1, out: accumulator.out, _value: accumulator._value + 1},
        identity: {in: 0, out: 0, _value: 0},
    )
    |> group()
    |> aggregateWindow(every: %s, fn: max, createEmpty: true)`, strconv.Quote(routerId), strconv.Quote(ospfName), startTime.Unix(), stopTime.Unix(), every))
	if err != nil {
		log.Printf("query fail:%v", err)
		return nil, ErrDatabase
	}

	return readDirectedTimeLinks(res)
}

func OSPFRouterLast10Tickers(ospfName string, routerId string) ([]bool, error) {
	if !Enabled {
		return nil, ErrDatabaseDisabled
	}
	res, err := allQuery.Query(context.Background(), fmt.Sprintf(`from(bucket: "network")
  |> range(start: -10m, stop: now())
  |> filter(fn: (r) => r["_measurement"] == "%s" and r["_field"] == "up" and (r.src == "%s" or r.dst == "%s"))
  |> drop(columns: ["dst","src"])
  |> aggregateWindow(every: 1m, fn: max, createEmpty: true)`, ospfName, routerId, routerId))
	if err != nil {
		log.Printf("query fail: %v", err)
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

func BatchRecordOSPF(asn uint32, ospf *entity.OSPF, t time.Time) error {
	if !Enabled {
		return ErrDatabaseDisabled
	}
	var points []*write.Point
	peerRouter := make(map[string]map[string]bool)
	for _, area := range *ospf {
		for _, link := range area.Links {
			points = append(points, influxdb2.NewPointWithMeasurement(fmt.Sprintf("ospf-%d", asn)).
				AddField("up", 1).
				AddTag("src", link.Src).
				AddTag("dst", link.Dst).
				SetTime(t))
			if peerRouter[link.Src] == nil {
				peerRouter[link.Src] = make(map[string]bool)
			}
			peerRouter[link.Src][link.Dst] = true
		}
	}
	if err := networkWrite.WritePoint(context.Background(), points...); err != nil {
		log.Printf("write record fail:%v", err)
		return ErrDatabase
	}

	var outDegree []*write.Point
	for router, peers := range peerRouter {
		outDegree = append(outDegree, influxdb2.NewPointWithMeasurement(fmt.Sprintf("ospf-%d", asn)).
			AddField("degree", len(peers)).
			AddTag("routerId", router).
			SetTime(t))
	}
	if err := peerCountWrite.WritePoint(context.Background(), points...); err != nil {
		log.Printf("write peer count fail:%v", err)
		return ErrDatabase
	}

	return nil
}
