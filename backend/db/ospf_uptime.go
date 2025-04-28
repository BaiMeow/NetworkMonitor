package db

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/consts"
	"log"
	"time"
)

func OSPFLinks(ospfName string, routerId string, startTime, stopTime time.Time, window time.Duration) ([]consts.LinkTime, error) {
	if !Enabled {
		return nil, ErrDatabaseDisabled
	}
	every, err := ParseWindow(window)
	if err != nil {
		return nil, err
	}

	res, err := dbQuery.Query(context.Background(), fmt.Sprintf(`from(bucket: "network")
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

	return ReadTimeLinks(res)
}
