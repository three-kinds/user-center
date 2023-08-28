package models

import "time"

type ResetPasswordCode struct {
	ID         int64     `gorm:"autoIncrement;primary_key"`
	Key        string    `gorm:"type:varchar(64);not null;uniqueIndex"`
	UserID     int64     `gorm:"not null;index"`
	Expiration time.Time `gorm:"not null;index"`
}
