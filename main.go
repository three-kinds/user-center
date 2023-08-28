package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/three-kinds/user-center/daos/models"
	"github.com/three-kinds/user-center/di"
	"github.com/three-kinds/user-center/initializers"
	"log"
	"net/http"
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
	server.Use(gin.Recovery())
	// build router
	router := server.Group("/api")
	router.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	di.RegisterAuthControllerRouter(router)

	log.Fatal(server.Run(fmt.Sprintf(":%d", initializers.Config.ServerPort)))
}
