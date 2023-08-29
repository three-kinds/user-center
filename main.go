package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/daos/models"
	"github.com/three-kinds/user-center/di"
	"github.com/three-kinds/user-center/initializers"
	"github.com/three-kinds/user-center/utils/frame_utils/middlewares"
	"log"
)

func init() {
	initializers.InitConfig("")
	initializers.InitDB(initializers.Config, &models.User{})
	initializers.InitLogger()
	initializers.InitSnowflakeNode(initializers.Config)
	initializers.InitValidators()
}

func main() {
	server := gin.New()
	server.Use(middlewares.JsonLogger())
	server.Use(gin.Recovery())
	server.MaxMultipartMemory = 64 << 20
	// build router
	router := server.Group("/api")
	di.RegisterOperationRouter(router)
	di.RegisterAuthControllerRouter(router)
	di.RegisterProfileControllerRouter(router)
	di.RegisterUserManagementControllerRouter(router)
	// run server
	log.Fatal(server.Run(fmt.Sprintf(":%d", initializers.Config.ServerPort)))
}
