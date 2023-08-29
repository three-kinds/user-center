package daos

import (
	"github.com/three-kinds/user-center/services/bo"
	"time"
)

type IUserDAO interface {
	CreateUser(user *bo.CreateUserBO, id int64, dateJoined time.Time) (*bo.UserBO, error)
	UpdatePassword(id int64, password string) error
	UpdateUser(id int64, updateUserBO *bo.UpdateUserBO) error
	UpdateProfile(id int64, updateUserBO *bo.UpdateProfileBO) error
	ListUsers(page int, size int, isActive *bool, isSuperuser *bool) ([]*bo.UserBO, error)
	Count(isActive *bool, isSuperuser *bool) (int64, error)
	GetUserByID(id int64) (*bo.UserBO, error)
	GetUserByUsername(username string) (*bo.UserBO, error)
	GetUserByEmail(email string) (*bo.UserBO, error)
	GetUserByPhoneNumber(number string) (*bo.UserBO, error)
	DeleteUserByID(id int64) error
}
