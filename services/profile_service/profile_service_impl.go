package profile_service

import (
	"github.com/three-kinds/user-center/daos"
	"github.com/three-kinds/user-center/services/bo"
	"github.com/three-kinds/user-center/utils/service_utils/password_utils"
	"github.com/three-kinds/user-center/utils/service_utils/se"
)

type ProfileServiceImpl struct {
	userDAO daos.IUserDAO
}

func (s *ProfileServiceImpl) PartialUpdateProfile(id int64, updateProfileBO *bo.UpdateProfileBO) (*bo.UserBO, error) {
	err := s.userDAO.UpdateProfile(id, updateProfileBO)
	if err != nil {
		return nil, err
	}
	return s.userDAO.GetUserByID(id)
}

func (s *ProfileServiceImpl) UpdatePassword(id int64, oldPassword string, newPassword string) error {
	user, err := s.userDAO.GetUserByID(id)
	if err != nil {
		return err
	}

	ok := password_utils.IsSamePassword(oldPassword, user.Password)
	if !ok {
		return se.ValidationError("raw password mismatch")
	}

	hashedPassword, err := password_utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	return s.userDAO.UpdatePassword(id, hashedPassword)
}

func NewProfileServiceImpl(userDAO daos.IUserDAO) *ProfileServiceImpl {
	return &ProfileServiceImpl{userDAO: userDAO}
}
