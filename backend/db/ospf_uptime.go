package db

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/consts"
	"log"
	"slices"
	"strconv"
	"time"
)

func OSPFLinks(ospfName string, routerId string, startTime, stopTime time.Time, window time.Duration) ([]consts.DirectedLinkTime, error) {
	if !Enabled {
		return nil, ErrDatabaseDisabled
	}
	every, err := ParseWindow(window)
	if err != nil {
		return nil, err
	}

	res, err := dbQuery.Query(context.Background(), fmt.Sprintf(`routerId = %s
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
	res, err := dbQuery.Query(context.Background(), fmt.Sprintf(`from(bucket: "network")
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
