package user_management_service

import "github.com/three-kinds/user-center/services/user_service"

type IUserManagementService interface {
	ListUsers(page int, size int, isActive *bool, isSuperuser *bool) (total int64, userList []*user_service.UserBO, err error)
	CreateUser(createUserBO *CreateUserBO) (*user_service.UserBO, error)
	GetUserByID(id int64) (*user_service.UserBO, error)
	GetUserByUsername(username string) (*user_service.UserBO, error)
	PartialUpdateUser(id int64, updateUserBO *UpdateUserBO) error
	DeleteUser(id int64) error
}
