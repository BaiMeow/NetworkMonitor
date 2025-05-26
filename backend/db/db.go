package db

import (
	"context"
	"fmt"
	"github.com/BaiMeow/NetworkMonitor/conf"
	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"log"
	"slices"
	"time"
)

var (
	networkWrite   api.WriteAPIBlocking
	peerCountWrite api.WriteAPIBlocking
	allQuery       api.QueryAPI
	Enabled        = false
)

var ErrDatabaseDisabled = fmt.Errorf("database not enabled")
var ErrInvalidInput = fmt.Errorf("input parameter invalid")

const (
	bucketNetwork   = "network"
	bucketPeerCount = "peerCount"
)

func Init() error {
	if conf.Influxdb.Addr == "" {
		return ErrDatabaseDisabled
	}
	c := influxdb.NewClient(conf.Influxdb.Addr, conf.Influxdb.Token)
	// normally less than 20 buckets, no check page
	buckets, err := c.BucketsAPI().FindBucketsByOrgName(context.Background(), conf.Influxdb.Org)
	if err != nil {
		return fmt.Errorf("find bucket fail:%v", err)
	}

	if slices.IndexFunc(*buckets, func(bucket domain.Bucket) bool {
		return bucket.Name == bucketNetwork
	}) == -1 {
		log.Printf("create bucket %s\n", bucketNetwork)
		org, err := c.OrganizationsAPI().FindOrganizationByName(context.Background(), conf.Influxdb.Org)
		if err != nil {
			return fmt.Errorf("org %s not existed:%v\n", conf.Influxdb.Org, err)
		}
		if _, err := c.BucketsAPI().CreateBucketWithName(context.Background(), org, bucketNetwork, domain.RetentionRule{
			EverySeconds: int64(conf.Uptime.StoreDuration / time.Second),
		}); err != nil {
			return fmt.Errorf("create bucket fail:%v", err)
		}
	}

	networkWrite = c.WriteAPIBlocking(conf.Influxdb.Org, bucketNetwork)
	peerCountWrite = c.WriteAPIBlocking(conf.Influxdb.Org, bucketPeerCount)
	allQuery = c.QueryAPI(conf.Influxdb.Org)

	Enabled = true

	return nil
}

var ErrDatabase = fmt.Errorf("database error")
