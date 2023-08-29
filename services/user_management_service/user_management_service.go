package user_management_service

import (
	"github.com/three-kinds/user-center/services/bo"
)

type IUserManagementService interface {
	ListUsers(page int, size int, isActive *bool, isSuperuser *bool) (total int64, userList []*bo.UserBO, err error)
	CreateUser(createUserBO *CreateUserBO) (*bo.UserBO, error)
	GetUserByID(id int64) (*bo.UserBO, error)
	GetUserByUsername(username string) (*bo.UserBO, error)
	PartialUpdateUser(id int64, updateUserBO *UpdateUserBO) error
	DeleteUser(id int64) error
}
