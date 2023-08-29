package vo

import "time"

type UserVO struct {
	ID          int64      `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Nickname    *string    `json:"nickname"`
	PhoneNumber *string    `json:"phone_number"`
	Avatar      *string    `json:"avatar"`
	DateJoined  time.Time  `json:"date_joined"`
	LastLogin   *time.Time `json:"last_login"`
	IsActive    bool       `json:"is_active"`
	IsSuperuser bool       `json:"is_superuser"`
}
