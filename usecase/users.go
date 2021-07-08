package usecase

import (
	"errors"
	"fmt"

	"bitbucket.org/bridce/fgp-go-boilerplate/models"
	"bitbucket.org/bridce/fgp-go-boilerplate/repo"
)

type UsersUsecaseStruct struct {
	users repo.UsersInterface
}

type UsersInterface interface {
	CreateUsers(data *models.RequestUser) (*models.ResponseUser, error)
}

func CreateGetTokenUsecaseImp(users repo.UsersInterface) UsersInterface {
	return &UsersUsecaseStruct{users}
}

func (g *UsersUsecaseStruct) CreateUsers(data *models.RequestUser) (*models.ResponseUser, error) {

	res, err := g.users.CreateUsers(*data)
	if err != nil {
		fmt.Println("[UsersUsecase.CreateUsers] error create")
		return nil, errors.New("error create")
	}

	return &models.ResponseUser{
		Nama:      res.Nama,
		Pekerjaan: res.Pekerjaan,
		ID:        res.ID,
		CreatedAt: res.CreatedAt,
	}, nil
}
