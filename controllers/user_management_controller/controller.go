package user_management_controller

import (
	"github.com/three-kinds/user-center/services/user_management_service"
)

type UserManagementController struct {
	userManagementService user_management_service.IUserManagementService
}

func NewUserManagementController(userManagementService user_management_service.IUserManagementService) *UserManagementController {
	return &UserManagementController{userManagementService: userManagementService}
}
