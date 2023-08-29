package daos

import (
	"github.com/three-kinds/user-center/services/bo"
	"github.com/three-kinds/user-center/services/user_management_service"
	"time"
)

type IUserDAO interface {
	CreateUser(user *user_management_service.CreateUserBO, id int64, dateJoined time.Time) (*bo.UserBO, error)
	UpdatePassword(id int64, password string) error
	UpdateUser(id int64, updateUserBO *user_management_service.UpdateUserBO) error
	ListUsers(page int, size int, isActive *bool, isSuperuser *bool) ([]*bo.UserBO, error)
	Count(isActive *bool, isSuperuser *bool) (int64, error)
	GetUserByID(id int64) (*bo.UserBO, error)
	GetUserByUsername(username string) (*bo.UserBO, error)
	GetUserByEmail(email string) (*bo.UserBO, error)
	GetUserByPhoneNumber(number string) (*bo.UserBO, error)
	DeleteUserByID(id int64) error
}
