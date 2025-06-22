package uptime

import (
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	"github.com/BaiMeow/NetworkMonitor/db"
	"github.com/jellydator/ttlcache/v3"
	"golang.org/x/sync/singleflight"
	"log"
	"time"
)

var cacheAllASN = ttlcache.New[string, map[uint32]struct{}](
	ttlcache.WithTTL[string, map[uint32]struct{}](time.Hour*12),
	ttlcache.WithDisableTouchOnHit[string, map[uint32]struct{}](),
	ttlcache.WithLoader(
		ttlcache.NewSuppressedLoader[string, map[uint32]struct{}](
			&ASNListLoader{}, &singleflight.Group{},
		)),
)

func init() {
	go cacheAllASN.Start()
}

type ASNListLoader struct {
}

var _ ttlcache.Loader[string, map[uint32]struct{}] = (*ASNListLoader)(nil)

func (l *ASNListLoader) Load(c *ttlcache.Cache[string, map[uint32]struct{}], key string) *ttlcache.Item[string, map[uint32]struct{}] {
	ASNs, err := db.AllASRecordAfter(fmt.Sprintf("bgp-%s", key), time.Now().Add(-conf.Uptime.StoreDuration))
	if err != nil {
		log.Printf("load recorded ASN list: %v", err)
		return nil
	}
	ASNSet := make(map[uint32]struct{})
	for _, ASN := range ASNs {
		ASNSet[ASN] = struct{}{}
	}
	return c.Set(key, ASNSet, ttlcache.DefaultTTL)
}
