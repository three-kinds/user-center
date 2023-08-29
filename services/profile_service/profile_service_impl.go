package profile_service

import (
	"github.com/three-kinds/user-center/services/bo"
	"github.com/three-kinds/user-center/services/user_service"
)

type ProfileServiceImpl struct {
	userService user_service.IUserService
}

func (s *ProfileServiceImpl) PartialUpdateProfile(id int64, updateProfileBO *UpdateProfileBO) (*bo.UserBO, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProfileServiceImpl) UpdatePassword(id int64, oldPassword string, newPassword string) error {
	//TODO implement me
	panic("implement me")
}

func NewProfileServiceImpl(userService user_service.IUserService) *ProfileServiceImpl {
	return &ProfileServiceImpl{userService: userService}
}
