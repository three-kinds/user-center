package main

import (
	"fmt"
	"github.com/three-kinds/user-center/daos/models"
	"github.com/three-kinds/user-center/di"
	"github.com/three-kinds/user-center/initializers"
	"github.com/three-kinds/user-center/services/bo"
	"github.com/three-kinds/user-center/services/user_management_service"
	"log"
)

var userList = []bo.CreateUserBO{
	{
		Email:       "admin@xx.com",
		Username:    "admin",
		Password:    "6uDU7xz5tY1Js",
		IsSuperuser: true,
	},
	{
		Email:       "ordinary@xx.com",
		Username:    "ordinary",
		Password:    "PRif6DRFtMT1",
		IsSuperuser: false,
	},
}

func init() {
	initializers.InitConfig("")
	initializers.InitDB(initializers.Config, &models.User{})
	initializers.InitSnowflakeNode(initializers.Config)
}

func mustCreateUser(s user_management_service.IUserManagementService, user *bo.CreateUserBO) *bo.UserBO {
	newUser, err := s.CreateUser(user)
	if err != nil {
		log.Panicln("create user failed", err)
	}
	return newUser
}

func main() {
	total := len(userList)
	userManagementService := di.NewUserManagementService()

	for i, user := range userList {
		oldUser, err := userManagementService.GetUserByUsername(user.Username)
		if err == nil {
			fmt.Printf("[%d/%d][pass] user %s has existed，id: %d \n", i+1, total, oldUser.Username, oldUser.ID)
			continue
		}

		newUser := mustCreateUser(userManagementService, &user)
		fmt.Printf("[%d/%d][created] create user %s success， id: %d \n", i+1, total, user.Username, newUser.ID)
	}
}
