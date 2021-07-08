package usecase

import (
	"fmt"

	"github.com/bridce/fgp-go-boilerplate/models"
	"github.com/bridce/fgp-go-boilerplate/repo"
	e "github.com/bridce/fgp-go-boilerplate/utils"
)

//your business logic can define here
type UsersUsecaseStruct struct {
	users repo.UsersInterface
}

type UsersInterface interface {
	CreateUsers(data *models.RequestUser) (*models.ResponseUser, *e.APIError)
}

func CreateGetTokenUsecaseImp(users repo.UsersInterface) UsersInterface {
	return &UsersUsecaseStruct{users}
}

func (g *UsersUsecaseStruct) CreateUsers(data *models.RequestUser) (*models.ResponseUser, *e.APIError) {

	res, err := g.users.CreateUsers(*data)
	if err != nil {
		fmt.Println("[UsersUsecase.CreateUsers] error create")
		return nil, &e.ErrCreateUsers
	}

	return &models.ResponseUser{
		Nama:      res.Nama,
		Pekerjaan: res.Pekerjaan,
		ID:        res.ID,
		CreatedAt: res.CreatedAt,
	}, nil
}
