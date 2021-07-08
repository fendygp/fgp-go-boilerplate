package controller

import (
	"fmt"
	"net/http"

	"bitbucket.org/bridce/fgp-go-boilerplate/models"
	"bitbucket.org/bridce/fgp-go-boilerplate/usecase"
	"bitbucket.org/bridce/fgp-go-boilerplate/utils"
	"github.com/gin-gonic/gin"
)

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
		utils.ErrorMessage(c, http.StatusBadRequest, "0001", "cannot read input data")
		fmt.Printf("[AuthControllerStruct.GetAccessToken] Error when bind json : %v\n", err)
		return
	}

	test, err := g.usersUsecase.CreateUsers(&data)
	if err != nil {
		utils.ErrorMessage(c, http.StatusBadRequest, "0002", err.Error())
	}
	utils.SuccessData(c, http.StatusOK, test)
}
