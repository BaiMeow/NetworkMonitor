package uptime

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/db"
	"github.com/BaiMeow/NetworkMonitor/graph"
	"log"
	"time"
)

func Init() {
	go tickerInsertDB()
}

func tickerInsertDB() {
	tk := make(chan time.Time)
	nextTrigger := time.Now()
	nextTrigger = nextTrigger.Add(-time.Duration(nextTrigger.Second()) * time.Second)
	nextTrigger = nextTrigger.Add(-time.Duration(nextTrigger.Nanosecond()))
	nextTrigger.Add(conf.Uptime.Interval)
	go func() {
		for {
			time.Sleep(time.Until(nextTrigger))
			tk <- nextTrigger
			nextTrigger = nextTrigger.Add(conf.Uptime.Interval)
		}
	}()
	for {
		now := <-tk
		gr := graph.BGP
		mp := make(map[uint32]int, len(gr.AS))
		for _, as := range gr.AS {
			mp[as.ASN] = 0
		}
		for _, lk := range gr.Link {
			mp[lk.Src]++
			mp[lk.Dst]++
		}
		err := db.BatchRecordASUp(mp, now)
		log.Println("record as up")
		if err != nil {
			log.Println(fmt.Errorf("record as up fail:%v", err))
		}
	}
}
