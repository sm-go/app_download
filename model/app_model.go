package model

import "time"

type App struct {
	Id          uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"column:title" json:"title"`
	Version     string    `gorm:"column:version" json:"version"`
	Description string    `gorm:"column:description" json:"description"`
	Size        string    `gorm:"column:size" json:"size"`
	OsType      string    `gorm:"column:os_type" json:"os_type"`
	DomainId    uint64    `gorm:"column:domain_id" json:"domain_id"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
