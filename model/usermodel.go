package model

import "time"

type User struct {
	Id        uint64    `gorm:"column:id;primaryKey" json:"id"`
	Username  string    `gorm:"column:username;notNull" json:"username"`
	Password  string    `gorm:"column:password;not null" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
