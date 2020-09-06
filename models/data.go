package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

const ISO8061Format = "%04d-%02d-%02dT%02d:%02d:%02d+%s"

type Data struct {
	UUID4         string `gorm:"uuid;primaryKey"`
	Title         string `gorm:"title;unique"`
	Timestamp     string `gorm:"-"`
	UnixTimestamp int64  `gorm:"unix_timestamp" json:"-"`
}

func (u *Data) AfterFind(_ *gorm.DB) (err error) {
	t := time.Unix(u.UnixTimestamp, 0).UTC()
	// tz offset is 00:00 because is always UTC (GMT)
	u.Timestamp = fmt.Sprintf(ISO8061Format, t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), "00:00")
	return
}
