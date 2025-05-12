package uptime

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/consts"
	"github.com/BaiMeow/NetworkMonitor/db"
	"github.com/BaiMeow/NetworkMonitor/utils"
	"time"
)

func AllASNRecord(bgpName string) ([]uint32, error) {
	if !db.Enabled {
		return nil, nil
	}
	ASNs, err := db.AllASRecordAfter(fmt.Sprintf("bgp-%s", bgpName), time.Now().Add(-conf.Uptime.StoreDuration))
	if err != nil {
		return nil, fmt.Errorf("get all recorded as fail:%v", err)
	}
	return ASNs, nil
}

func Last10BGPTickerRecord(bgpName string, asn uint32) ([]bool, error) {
	if !db.Enabled {
		return nil, nil
	}
	records, err := db.BGPASNLast10Tickers(fmt.Sprintf("bgp-%s", bgpName), asn)
	if err != nil {
		return nil, fmt.Errorf("get last 10 tickers fail:%v", err)
	}
	return records, nil
}

func BGPLinks(bgpName string, asn uint32, window, t time.Duration) ([]consts.LinkTime, error) {
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
	links, err := db.BGPLinks(fmt.Sprintf("bgp-%s", bgpName), asn, startTime, stopTime, window)
	if err != nil {
		return links, fmt.Errorf("get links fail:%v", err)
	}
	return links, nil
}

func OSPFLinks(asn uint32, routerId string, window, t time.Duration) (*[2][]consts.LinkTime, error) {
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
	links, err := db.OSPFLinks(fmt.Sprintf("ospf-%d", asn), routerId, startTime, stopTime, window)
	if err != nil {
		return links, fmt.Errorf("get links fail:%v", err)
	}
	return links, nil
}

func Last10OSPFTickerRecord(asn uint32, routerId string) ([]bool, error) {
	if !db.Enabled {
		return nil, nil
	}
	records, err := db.OSPFRouterLast10Tickers(fmt.Sprintf("ospf-%d", asn), routerId)
	if err != nil {
		return nil, fmt.Errorf("get last 10 tickers fail:%v", err)
	}
	return records, nil
}
