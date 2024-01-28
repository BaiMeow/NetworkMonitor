package utils

import (
	"github.com/BaiMeow/NetworkMonitor/conf"
	"time"
)

// WithTimeout run f with timeout, return if timeout
func WithTimeout(f func()) bool {
	ok := make(chan struct{})
	go func() {
		f()
		close(ok)
	}()
	select {
	case <-ok:
		return false
	case <-time.After(time.Second * time.Duration(conf.Interval*4/5)):
		return true
	}
}
