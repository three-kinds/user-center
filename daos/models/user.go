package models

import (
	"time"
)

type User struct {
	ID          int64      `gorm:"primary_key"`
	Username    string     `gorm:"type:varchar(64);not null;uniqueIndex"`
	Email       string     `gorm:"type:varchar(64);not null;uniqueIndex"`
	Password    string     `gorm:"type:varchar(255);not null"`
	Nickname    *string    `gorm:"type:varchar(64);index"`
	PhoneNumber *string    `gorm:"type:varchar(32);uniqueIndex"`
	Avatar      *string    `gorm:"type:varchar(255)"`
	DateJoined  time.Time  `gorm:"not null;index"`
	LastLogin   *time.Time `gorm:"index"`
	IsActive    bool       `gorm:"not null;index"`
	IsSuperuser bool       `gorm:"not null;index"`
}
