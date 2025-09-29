package entity

import "time"

type User struct {
	ID           int64
	Username     string
	PasswordHash string
}

type RefreshToken struct {
	ID        int64
	UserID    int64
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}
