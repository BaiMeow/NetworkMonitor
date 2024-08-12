package db

import (
	"github.com/BaiMeow/NetworkMonitor/utils"
	"log"
	"time"
)

type ASUp struct {
	ASN       uint32
	CreatedAt time.Time `gorm:"index"`
}

func BatchRecordASUp(ASNs []uint32, t time.Time) error {
	ups := utils.Map(ASNs, func(asn uint32) *ASUp { return &ASUp{ASN: asn, CreatedAt: t} })
	if tx := db.Create(ups); tx.Error != nil {
		log.Printf("record as up fail:%v\n", tx.Error)
		return ErrDatabase
	}
	return nil
}

func AllASRecordAfter(after time.Time) ([]uint32, error) {
	var asns []uint32
	if tx := db.Select("asn").Distinct("asn").Where("create_at > ?", after).Find(&asns); tx.Error != nil {
		log.Printf("get all recorded as fail:%v\n", tx.Error)
		return nil, ErrDatabase
	}
	return asns, nil
}

func CleanASUptimeBefore(before time.Time) error {
	if tx := db.Delete(&ASUp{}, "created_at < ?", before); tx.Error != nil {
		log.Printf("clean by time fail:%v\n", tx.Error)
		return ErrDatabase
	}
	return nil
}

func BGPASNLast10Tickers(asn uint32, last time.Time) ([]time.Time, error) {
	var t []time.Time
	if tx := db.Select("created_at").
		Where("asn = ?", asn).
		Where("created_at <= ? AND created_at => ?", last, utils.TickOffset(last, -9)).
		Order("created_at").
		Limit(10).
		Find(&t); tx.Error != nil {
		log.Printf("get last 10 tickers fail:%v\n", tx.Error)
		return nil, ErrDatabase
	}
	return t, nil
}
