package router

import (
	"fmt"

	"github.com/fgp-go-boilerplate/controller"
	"github.com/fgp-go-boilerplate/repo"
	"github.com/fgp-go-boilerplate/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	usersRepo := repo.UsersTokenRepoImp()

	usersUsecase := usecase.CreateGetTokenUsecaseImp(usersRepo)

	v1 := router.Group("api/v1")
	{
		newRoute := v1.Group("reqresin")

		controller.CreateUsersControllerImp(newRoute, usersUsecase)
	}

	fmt.Println("about to start the application...")

	return router
}
