package uptime

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/consts"
	"github.com/BaiMeow/NetworkMonitor/db"
	"github.com/BaiMeow/NetworkMonitor/utils"
	"log"
	"time"
)

func AllASNRecord() ([]uint32, error) {
	if !db.Enabled {
		return nil, nil
	}
	ASNs, err := db.AllASRecordAfter(time.Now().Add(-conf.Uptime.StoreDuration))
	if err != nil {
		return nil, fmt.Errorf("get all recorded as fail:%v", err)
	}
	return ASNs, nil
}

func Last10TickerRecord(asn uint32) ([]bool, error) {
	if !db.Enabled {
		return nil, nil
	}
	last := utils.LastUptimeTick()
	records, err := db.BGPASNLast10Tickers(asn, last)
	if err != nil {
		return nil, fmt.Errorf("get last 10 tickers fail:%v", err)
	}
	up := make([]bool, 10)
	for _, record := range records {
		offset := int(last.Sub(record) / conf.Uptime.Interval)
		if offset >= 10 || offset < 0 {
			log.Printf("record time wrong:%v, last:%v, offset:%v asn:%v", record, last, offset, asn)
			continue
		}
		up[offset] = true
	}
	return up, nil
}

func Links(asn uint32, window, t time.Duration) ([]consts.LinkTime, error) {
	if !db.Enabled {
		return nil, nil
	}
	if t > time.Hour*24 && window != time.Hour {
		return nil, fmt.Errorf("invalid window size %s for time range %s", window, t)
	}
	if t <= time.Hour*24 && window != time.Minute {
		return nil, fmt.Errorf("invalid window size %s for time range %s", window, t)
	}
	stopTime := utils.LastUptimeTick()
	startTime := stopTime.Add(-t)
	links, err := db.BGPLinks(asn, startTime, stopTime, window)
	if err != nil {
		return links, fmt.Errorf("get links fail:%v", err)
	}
	return links, nil
}
