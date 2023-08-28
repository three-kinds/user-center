package services

import "github.com/three-kinds/user-center/services/bo"

type IUserManagementService interface {
	ListUsers(page int, size int, isActive *bool, isSuperuser *bool) (total int64, userList []*bo.UserDisplayBO, err error)
	CreateUser(createUserBO *bo.CreateUserBO) (*bo.UserDisplayBO, error)
	GetUserByID(id int64) (*bo.UserDisplayBO, error)
	GetUserByUsername(username string) (*bo.UserDisplayBO, error)
	PartialUpdateUser(id int64, updateUserBO *bo.UpdateUserBO) error
	DeleteUser(id int64) error
}
