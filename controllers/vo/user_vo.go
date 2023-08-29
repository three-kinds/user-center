package vo

import (
	"github.com/three-kinds/user-center/services/bo"
	"time"
)

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

func TUserBO2UserVO(bo *bo.UserBO) *UserVO {
	return &UserVO{
		ID:          bo.ID,
		Username:    bo.Username,
		Email:       bo.Email,
		Nickname:    bo.Nickname,
		PhoneNumber: bo.PhoneNumber,
		Avatar:      bo.Avatar,
		DateJoined:  bo.DateJoined,
		LastLogin:   bo.LastLogin,
		IsActive:    bo.IsActive,
		IsSuperuser: bo.IsSuperuser,
	}
}
