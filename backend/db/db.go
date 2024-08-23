package db

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"log"
	"slices"
	"time"
)

var (
	bgpWrite api.WriteAPIBlocking
	bgpQuery api.QueryAPI
)

const bucketBGPUptime = "bgp-uptime"

func Init() error {
	c := influxdb2.NewClient(conf.Influxdb.Addr, conf.Influxdb.Token)
	// normally less than 20 buckets, no check page
	buckets, err := c.BucketsAPI().FindBucketsByOrgName(context.Background(), conf.Influxdb.Org)
	if err != nil {
		return fmt.Errorf("find bucket fail:%v", err)
	}

	if slices.IndexFunc(*buckets, func(bucket domain.Bucket) bool {
		return bucket.Name == bucketBGPUptime
	}) == -1 {
		log.Printf("create bucket %s\n", bucketBGPUptime)
		org, err := c.OrganizationsAPI().FindOrganizationByName(context.Background(), conf.Influxdb.Org)
		if err != nil {
			return fmt.Errorf("org %s not existed:%v\n", conf.Influxdb.Org, err)
		}
		if _, err := c.BucketsAPI().CreateBucketWithName(context.Background(), org, bucketBGPUptime, domain.RetentionRule{
			EverySeconds: int64(conf.Uptime.StoreDuration / time.Second),
		}); err != nil {
			return fmt.Errorf("create bucket fail:%v", err)
		}
	}

	bgpWrite = c.WriteAPIBlocking(conf.Influxdb.Org, bucketBGPUptime)
	bgpQuery = c.QueryAPI(conf.Influxdb.Org)

	return nil
}

var ErrDatabase = fmt.Errorf("database error")
