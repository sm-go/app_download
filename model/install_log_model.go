package model

import "time"

type InstallLog struct {
	Id        uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	IpAddress string    `gorm:"column:ip_address" json:"ip_address"`
	DeviceId  string    `gorm:"column:device_id" json:"device_id"`
	AppId     uint64    `gorm:"app_id" json:"app_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
