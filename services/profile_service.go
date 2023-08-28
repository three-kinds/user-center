package services

import (
	"github.com/three-kinds/user-center/services/bo"
)

type IProfileService interface {
	GetProfile(accessToken string) (*bo.UserDisplayBO, error)
	PartialUpdateProfile(updateProfileBO *bo.UpdateProfileBO) (*bo.UserDisplayBO, error)
	UpdatePassword(oldPassword string, newPassword string) error
}
