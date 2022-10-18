package model

import "time"

type DownloadLog struct {
	Id        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	IpAddress string    `gorm:"ip_address" json:"ip_address"`
	AppId     uint64    `gorm:"app_id" json:"app_id"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}
