package services

import (
	"github.com/three-kinds/user-center/daos"
	"github.com/three-kinds/user-center/initializers"
	"github.com/three-kinds/user-center/services/bo"
	"github.com/three-kinds/user-center/utils/service_utils/password_utils"
	"time"
)

type UserManagementServiceImpl struct {
	userDAO daos.IUserDAO
}

func (s *UserManagementServiceImpl) ListUsers(page int, size int, isActive *bool, isSuperuser *bool) (total int64, userList []*bo.UserDisplayBO, err error) {
	total, err = s.userDAO.Count(isActive, isSuperuser)
	if err != nil {
		return 0, nil, err
	}
	userList, err = s.userDAO.ListUsers(page, size, isActive, isSuperuser)
	if err != nil {
		return 0, nil, err
	}
	return total, userList, nil
}

func (s *UserManagementServiceImpl) CreateUser(createUserBO *bo.CreateUserBO) (*bo.UserDisplayBO, error) {
	id := int64(initializers.SnowflakeNode.Generate())
	dateJoined := time.Now()
	hashPassword, err := password_utils.HashPassword(createUserBO.Password)
	if err != nil {
		return nil, err
	}
	createUserBO.Password = hashPassword

	return s.userDAO.CreateUser(createUserBO, id, dateJoined)
}

func (s *UserManagementServiceImpl) PartialUpdateUser(id int64, updateUserBO *bo.UpdateUserBO) error {
	if updateUserBO.Password != nil {
		hashPassword, err := password_utils.HashPassword(*updateUserBO.Password)
		if err != nil {
			return err
		}
		updateUserBO.Password = &hashPassword
	}

	return s.userDAO.UpdateUser(id, updateUserBO)
}

func (s *UserManagementServiceImpl) GetUserByID(id int64) (*bo.UserDisplayBO, error) {
	return s.userDAO.GetUserByID(id)
}

func (s *UserManagementServiceImpl) GetUserByUsername(username string) (*bo.UserDisplayBO, error) {
	return s.userDAO.GetUserByUsername(username)
}

func (s *UserManagementServiceImpl) DeleteUser(id int64) error {
	return s.userDAO.DeleteUserByID(id)
}

func NewUserManagementServiceImpl(userDAO daos.IUserDAO) *UserManagementServiceImpl {
	return &UserManagementServiceImpl{userDAO: userDAO}
}
