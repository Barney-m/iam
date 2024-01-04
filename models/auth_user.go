package models

import "time"

type AuthUser struct {
	UserId      string `gorm:"primaryKey"`
	FullName    string
	Email       string
	Password    string
	MobileNo    string
	Address     string
	Dob         time.Time
	Gender      string
	ProfilePic  string
	IsActive    int8
	IsFirstTime int8
	IsDeleted   int8
	BaseInfo
}
