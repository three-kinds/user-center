package profile_controller

import (
	"github.com/three-kinds/user-center/services/profile_service"
)

type ProfileController struct {
	profileService profile_service.IProfileService
}

func NewProfileController(profileService profile_service.IProfileService) *ProfileController {
	return &ProfileController{profileService: profileService}
}
