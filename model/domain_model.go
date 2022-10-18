package model

import "time"

type Domain struct {
	Id          uint64    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	DownloadUrl string    `gorm:"column:download_url" json:"download_url"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
