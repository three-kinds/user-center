package user_management_service

import (
	"github.com/three-kinds/user-center/services/bo"
)

type IUserManagementService interface {
	ListUsers(page int, size int, isActive *bool, isSuperuser *bool) (total int64, userList []*bo.UserBO, err error)
	CreateUser(createUserBO *bo.CreateUserBO) (*bo.UserBO, error)
	GetUserByID(id int64) (*bo.UserBO, error)
	PartialUpdateUser(id int64, updateUserBO *bo.UpdateUserBO) (*bo.UserBO, error)
	GetUserByUsername(username string) (*bo.UserBO, error)
	DeleteUser(id int64) error
}
