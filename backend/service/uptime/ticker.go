package uptime

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/db"
	"github.com/BaiMeow/NetworkMonitor/graph"
	"github.com/BaiMeow/NetworkMonitor/parse"
	"github.com/BaiMeow/NetworkMonitor/utils"
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
		err := db.BatchRecordASUp(utils.Map(graph.BGP.AS, func(as *parse.AS) uint32 { return as.ASN }), now)
		log.Println("record as up")
		if err != nil {
			log.Println(fmt.Errorf("record as up fail:%v", err))
		}
	}
}
