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
		bgps := graph.GetAllBGP()
		for k, gr := range bgps {
			data, t := gr.GetData()
			// if data too old, skip
			if t.Add(conf.Uptime.Interval).Before(now) {
				continue
			}
			err := db.BatchRecordBGP(k, data, t)
			log.Printf("record bgp graph %s at %v", k, now)
			if err != nil {
				log.Println(fmt.Errorf("record %s fail:%v", k, err))
			}
		}
		ospfs := graph.GetAllOSPF()
		for k, gr := range ospfs {
			data, t := gr.GetData()
			// if data too old, skip
			if t.Add(conf.Uptime.Interval).Before(now) {
				continue
			}
			err := db.BatchRecordOSPF(k, data, t)
			log.Printf("record ospf graph %d at %v", k, now)
			if err != nil {
				log.Println(fmt.Errorf("record %d fail:%v", k, err))
			}
		}
	}
}
