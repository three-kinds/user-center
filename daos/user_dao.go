package daos

import (
	"github.com/three-kinds/user-center/services/user_management_service"
	"github.com/three-kinds/user-center/services/user_service"
	"time"
)

type IUserDAO interface {
	CreateUser(user *user_management_service.CreateUserBO, id int64, dateJoined time.Time) (*user_service.UserBO, error)
	UpdatePassword(id int64, password string) error
	UpdateUser(id int64, updateUserBO *user_management_service.UpdateUserBO) error
	ListUsers(page int, size int, isActive *bool, isSuperuser *bool) ([]*user_service.UserBO, error)
	Count(isActive *bool, isSuperuser *bool) (int64, error)
	GetUserByID(id int64) (*user_service.UserBO, error)
	GetUserByUsername(username string) (*user_service.UserBO, error)
	GetUserByEmail(email string) (*user_service.UserBO, error)
	GetUserByPhoneNumber(number string) (*user_service.UserBO, error)
	DeleteUserByID(id int64) error
}
