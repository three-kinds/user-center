package di

import (
	"github.com/three-kinds/user-center/services/auth_service"
	"github.com/three-kinds/user-center/services/profile_service"
	"github.com/three-kinds/user-center/services/user_management_service"
	"github.com/three-kinds/user-center/services/user_service"
)

// 单体

func NewUserService() user_service.IUserService {

}

// 复合

func NewUserManagementService() user_management_service.IUserManagementService {
	return user_management_service.NewUserManagementServiceImpl(NewUserDAO())
}

func NewAuthService() auth_service.IAuthService {
	return auth_service.NewAuthServiceImpl()
}

func NewProfileService() profile_service.IProfileService {
	return profile_service.NewProfileServiceImpl()
}
