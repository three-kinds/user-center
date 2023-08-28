package user_service

type IUserService interface {
	LoginByAccountAndPassword(account string, password string) (*UserBO, error)
	GetActiveUserByID(userID int64) (*UserBO, error)
	GetActiveUserByEmail(email string) (*UserBO, error)
	RegisterUserByEmailPassword(email string, password string) error
	ResetPassword(id int64, password string) error
}
