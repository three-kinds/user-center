package profile_service

import (
	"github.com/three-kinds/user-center/services/user_service"
)

type IProfileService interface {
	GetProfile(accessToken string) (*user_service.UserBO, error)
	PartialUpdateProfile(updateProfileBO *UpdateProfileBO) (*user_service.UserBO, error)
	UpdatePassword(oldPassword string, newPassword string) error
}
