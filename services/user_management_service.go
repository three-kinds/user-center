package services

import "github.com/three-kinds/user-center/services/bo"

type IUserManagementService interface {
	ListUsers(page int, size int, isActive *bool, isSuperuser *bool) (total int, userList []*bo.UserDisplayBO, err error)
	CreateUser(createUserBO *bo.CreateUserBO) error
	ReadUser(id int64) (*bo.UserDisplayBO, error)
	PartialUpdateUser(updateUserBO *bo.UpdateUserBO) error
	DeleteUser(id int64) error
}
