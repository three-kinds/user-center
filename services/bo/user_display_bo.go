package bo

import "time"

type UserDisplayBO struct {
	ID          int64
	Username    string
	Email       string
	Nickname    *string
	PhoneNumber *string
	Avatar      *string
	DateJoined  time.Time
	LastLogin   *time.Time
	IsActive    bool
	IsSuperuser bool
}
