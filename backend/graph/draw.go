package graph

import (
	"github.com/BaiMeow/NetworkMonitor/conf"
	"log"
	"strconv"
	"sync"
	"time"
)

type DrawerCleaner interface {
	Draw()
	CleanUp()
}

func draw() {
	var wg sync.WaitGroup
	fullLock.RLock()
	defer fullLock.RUnlock()
	drawSingle := func(name string, drawable DrawerCleaner) {
		defer wg.Done()
		alert := time.AfterFunc(conf.ProbeTimeout, func() {
			log.Printf("probe %s timeout but the goroutine is still running, a timeout should be added to the probe.\n", name)
		})
		drawable.Draw()
		t := time.Now()
		alert.Stop()
		dur := time.Since(t)
		if dur > time.Second*5 {
			log.Printf("probe %s slow draw: %v\n", name, dur)
		}
	}
	for name, v := range bgp {
		wg.Add(1)
		go drawSingle(name, v)
	}
	for name, v := range ospf {
		wg.Add(1)
		go drawSingle(strconv.FormatUint(uint64(name), 10), v)
	}

	select {
	case <-time.After(conf.ProbeTimeout):
	case <-func() <-chan struct{} {
		done := make(chan struct{})
		go func() {
			wg.Wait()
			close(done)
		}()
		return done
	}():
	}
}
