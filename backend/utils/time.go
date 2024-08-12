package utils

import (
	"github.com/BaiMeow/NetworkMonitor/conf"
	"time"
)

func LastUptimeTick() time.Time {
	t := time.Now().Add(-conf.Uptime.StoreDuration)
	t.Add(-time.Duration(t.Second()) * time.Second)
	t.Add(-time.Duration(t.Nanosecond()))
	return t
}

func TickOffset(t time.Time, offset int) time.Time {
	return t.Add(time.Duration(offset) * conf.Uptime.Interval)
}
