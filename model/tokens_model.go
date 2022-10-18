package model

import (
	"time"

	"github.com/google/uuid"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenData struct {
	SS        string
	ID        uuid.UUID
	ExpiresIn time.Duration
}
