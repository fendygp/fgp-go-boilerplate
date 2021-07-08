package controller

import (
	"fmt"
	"net/http"

	"github.com/fgp-go-boilerplate/models"
	"github.com/fgp-go-boilerplate/usecase"
	"github.com/fgp-go-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

//your controller to handle request define here
type UsersControllerStruct struct {
	usersUsecase usecase.UsersInterface
}

func CreateUsersControllerImp(router *gin.RouterGroup, usersUsecase usecase.UsersInterface) {
	inDB := UsersControllerStruct{usersUsecase}

	router.POST("/users", inDB.CreateUsers)
}

func (g *UsersControllerStruct) CreateUsers(c *gin.Context) {
	var data models.RequestUser

	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, &utils.ErrInput)
		fmt.Printf("[UsersControllerStruct.CreateUsers] Error when bind json : %v\n", err)
		return
	}

	test, err2 := g.usersUsecase.CreateUsers(&data)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, err2)
	}
	utils.SuccessData(c, http.StatusOK, test)
}
