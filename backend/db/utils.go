package db

import (
	"github.com/BaiMeow/NetworkMonitor/consts"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"log"
	"time"
)

func ParseWindow(window time.Duration) (string, error) {
	if window == time.Minute {
		return "1m", nil
	} else if window == time.Hour {
		return "1h", nil
	} else {
		log.Printf("invalid window:%v", window)
		return "", ErrInvalidInput
	}
}

func ReadTimeLinks(res *api.QueryTableResult) ([]consts.LinkTime, error) {
	var points []consts.LinkTime
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