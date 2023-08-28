package di

import (
	"github.com/three-kinds/user-center/daos"
	"github.com/three-kinds/user-center/initializers"
)

func NewUserDAO() daos.IUserDAO {
	return daos.NewUserDAOImpl(initializers.DB)
}

func NewCaptchaDAO() daos.ICaptchaDAO {
	return daos.NewCaptchaDAOImpl(initializers.DB)
}

func NewResetPasswordCodeDAO() daos.IResetPasswordCodeDAO {
	return daos.NewResetPasswordCodeDAOImpl(initializers.DB)
}
