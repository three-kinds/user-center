package models

import "time"

type Captcha struct {
	ID         int64     `gorm:"autoIncrement;primary_key"`
	Key        string    `gorm:"type:varchar(64);not null;uniqueIndex"`
	Answer     string    `gorm:"not null"`
	Expiration time.Time `gorm:"not null;index"`
}
