package db

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/consts"
	"log"
	"slices"
	"time"
)

func OSPFLinks(ospfName string, routerId string, startTime, stopTime time.Time, window time.Duration) (*[2][]consts.LinkTime, error) {
	var directedLinkTime [2][]consts.LinkTime
	if !Enabled {
		return nil, ErrDatabaseDisabled
	}
	every, err := ParseWindow(window)
	if err != nil {
		return nil, err
	}

	// out degree
	res1, err := dbQuery.Query(context.Background(), fmt.Sprintf(`from(bucket: "network")
  |> range(start: %d, stop: %d)
  |> filter(fn: (r) => r["_measurement"] == "%s" and r.src == "%s")
  |> group(columns: ["_time"])
  |> count(column: "_value")
  |> group()
  |> aggregateWindow(every: %s, fn: max, createEmpty: true)
  |> yield(name: "max")`, startTime.Unix(), stopTime.Unix(), ospfName, routerId, every))
	if err != nil {
		log.Printf("query fail:%v", err)
		return nil, ErrDatabase
	}

	directedLinkTime[0], err = ReadTimeLinks(res1)
	if err != nil {
		return nil, err
	}

	// in degree
	res2, err := dbQuery.Query(context.Background(), fmt.Sprintf(`from(bucket: "network")
  |> range(start: %d, stop: %d)
  |> filter(fn: (r) => r["_measurement"] == "%s" and r.dst=="%s")
  |> group(columns: ["_time"])
  |> count(column: "_value")
  |> group()
  |> aggregateWindow(every: %s, fn: max, createEmpty: true)
  |> yield(name: "max")`, startTime.Unix(), stopTime.Unix(), ospfName, routerId, every))
	if err != nil {
		log.Printf("query fail:%v", err)
		return nil, ErrDatabase

	}
	directedLinkTime[1], err = ReadTimeLinks(res2)
	if err!=nil{
		return nil,err
	}
	return &directedLinkTime, err
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
