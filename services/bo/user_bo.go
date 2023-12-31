package bo

import "time"

type UserBO struct {
	ID          int64
	Username    string
	Email       string
	Password    string
	Nickname    *string
	PhoneNumber *string
	Avatar      *string
	DateJoined  time.Time
	LastLogin   *time.Time
	IsActive    bool
	IsSuperuser bool
}
