package di

import (
	"github.com/three-kinds/user-center/daos"
	"github.com/three-kinds/user-center/initializers"
	"github.com/three-kinds/user-center/services"
)

func NewUserManagementService() services.IUserManagementService {
	return services.NewUserManagementServiceImpl(daos.NewUserDAOImpl(initializers.DB))
}
