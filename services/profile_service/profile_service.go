package profile_service

import (
	"github.com/three-kinds/user-center/services/bo"
)

type IProfileService interface {
	PartialUpdateProfile(id int64, updateProfileBO *bo.UpdateProfileBO) (*bo.UserBO, error)
	UpdatePassword(id int64, oldPassword string, newPassword string) error
}
