package dto

import "time"

type RequestUserLogin struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RequestUserRefreshToken struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" binding:"required"`
}

type ResponseUser struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	CreatedAt time.Time
}
