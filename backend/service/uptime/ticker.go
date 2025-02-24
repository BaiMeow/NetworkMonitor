package uptime

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/db"
	"github.com/BaiMeow/NetworkMonitor/graph"
	"log"
	"maps"
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
		grs:= graph.GetAllBGP()
		for gr :=range maps.Values(grs){
			data, t := gr.GetData()
			if t.Add(2 * conf.Interval).Before(now) {
				continue
			}
			mp := make(map[uint32]int, len(data.AS))
			for _, as := range data.AS {
				mp[as.ASN] = 0
			}
			for _, lk := range data.Link {
				mp[lk.Src]++
				mp[lk.Dst]++
			}
			//TODO: fix for multi bgp graph
			err := db.BatchRecordASUp(mp, t)
			log.Printf("record as %d links %d at %v", len(mp), len(data.Link), now)
			if err != nil {
				log.Println(fmt.Errorf("record as up fail:%v", err))
			}
		}
	}
}
