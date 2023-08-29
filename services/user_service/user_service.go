package user_service

import "github.com/three-kinds/user-center/services/bo"

type IUserService interface {
	GetUserByUsername(username string) (*bo.UserBO, error)
	LoginByAccountAndPassword(account string, password string) (*bo.UserBO, error)
	GetActiveUserByID(userID int64) (*bo.UserBO, error)
	GetActiveUserByEmail(email string) (*bo.UserBO, error)
	RegisterUserByEmailPassword(email string, password string) error
	ResetPassword(id int64, password string) error
}
