package daos

import (
	"github.com/three-kinds/user-center/services/bo"
	"time"
)

type IUserDAO interface {
	CreateUser(user *bo.CreateUserBO, id int64, dateJoined time.Time) (*bo.UserDisplayBO, error)
	UpdatePassword(id int64, password string) error
	UpdateUser(id int64, updateUserBO *bo.UpdateUserBO) error
	ListUsers(page int, size int, isActive *bool, isSuperuser *bool) ([]*bo.UserDisplayBO, error)
	Count(isActive *bool, isSuperuser *bool) (int64, error)
	GetUserByID(id int64) (*bo.UserDisplayBO, error)
	GetUserByUsername(username string) (*bo.UserDisplayBO, error)
	GetUserByEmail(email string) (*bo.UserDisplayBO, error)
	GetUserByPhoneNumber(number string) (*bo.UserDisplayBO, error)
	DeleteUserByID(id int64) error
}
