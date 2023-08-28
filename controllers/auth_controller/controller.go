package auth_controller

import (
	"github.com/three-kinds/user-center/services/auth_service"
)

type AuthController struct {
	authService auth_service.IAuthService
}

func NewAuthController(authService auth_service.IAuthService) *AuthController {
	return &AuthController{authService: authService}
}
