package uptime

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/db"
	"github.com/BaiMeow/NetworkMonitor/utils"
	"log"
	"time"
)

func AllASNRecord() ([]uint32, error) {
	if !db.Enabled{
		return nil, nil
	}
	ASNs, err := db.AllASRecordAfter(time.Now().Add(-conf.Uptime.StoreDuration))
	if err != nil {
		return nil, fmt.Errorf("get all recorded as fail:%v", err)
	}
	return ASNs, nil
}

func Last10TickerRecord(asn uint32) ([]bool, error) {
	if !db.Enabled{
		return nil, nil
	}
	last := utils.LastUptimeTick()
	records, err := db.BGPASNLast10Tickers(asn, last)
	if err != nil {
		return nil, fmt.Errorf("get last 10 tickers fail:%v", err)
	}
	up := make([]bool, 10)
	for _, record := range records {
		offset := last.Sub(record) / conf.Uptime.Interval
		if offset >= 10 || offset < 0 {
			log.Printf("record time wrong:%v, last:%v, offset:%v asn:%v", record, last, offset, asn)
			continue
		}
		up[offset] = true
	}
	return up, nil
}
