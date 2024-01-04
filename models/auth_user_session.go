package models

import "time"

type AuthUserSession struct {
	SessionId    uint64
	Email        string
	RefreshToken string
	UserAgent    string
	ClientIp     string
	IsBlocked    int8
	ExpiredAt    time.Time
	BaseInfo
}
