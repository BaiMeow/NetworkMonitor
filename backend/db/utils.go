package db

import (
	"github.com/BaiMeow/NetworkMonitor/consts"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/query"
	"log"
	"time"
)

func parseWindow(window time.Duration) (string, error) {
	if window == time.Minute {
		return "1m", nil
	} else if window == time.Hour {
		return "1h", nil
	} else {
		log.Printf("invalid window:%v", window)
		return "", ErrInvalidInput
	}
}

func readQueryTableResultFunc[T any](res *api.QueryTableResult, f func(rc *query.FluxRecord) T) ([]T, error) {
	var points []T
	for res.Next() {
		points = append(points, f(res.Record()))
	}
	return points, nil
}

func readTimeLinks(res *api.QueryTableResult) ([]consts.LinkTime, error) {
	return readQueryTableResultFunc(res, func(rc *query.FluxRecord) consts.LinkTime {
		v, ok := rc.Value().(int64)
		if !ok {
			log.Printf("convert influxdb value fail: %v", rc)
		}
		return consts.LinkTime{
			Time:  rc.Time(),
			Links: int(v),
		}
	})
}

func readDirectedTimeLinks(res *api.QueryTableResult) ([]consts.DirectedLinkTime, error) {
	return readQueryTableResultFunc(res, func(rc *query.FluxRecord) consts.DirectedLinkTime {
		in, _ := rc.ValueByKey("in").(int64)
		out, _ := rc.ValueByKey("out").(int64)
		return consts.DirectedLinkTime{
			Time:      rc.Time(),
			InDegree:  int(in),
			OutDegree: int(out),
		}
	})
}
