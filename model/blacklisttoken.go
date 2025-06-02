package model

import "time"

type BlacklistedToken struct {
	ID        uint
	Token     string
	ExpiresAt time.Time
	CreatedAT time.Time
}