package repo

import (
	"crypto/tls"
	"encoding/json"
	"errors"

	"github.com/fgp-go-boilerplate/models"
	"github.com/go-resty/resty/v2"
)

//your function for query db, consume another service cant define here
type UsersRepoStruct struct {
}

type UsersInterface interface {
	CreateUsers(data models.RequestUser) (*models.ResponseUser, error)
}

func UsersTokenRepoImp() UsersInterface {
	return &UsersRepoStruct{}
}

func (g *UsersRepoStruct) CreateUsers(data models.RequestUser) (*models.ResponseUser, error) {
	client := resty.New()
	client.SetDebug(true)
	client.SetContentLength(true)
	url := "https://reqres.in/api/users"

	resUser := &models.ResponseUserAPI{}
	resp, err := client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetRedirectPolicy(resty.FlexibleRedirectPolicy(15)).R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"name":"` + data.Nama + `", "job":"` + data.Pekerjaan + `"}`).
		SetResult(resUser).
		Post(url)

	if err != nil {
		return nil, errors.New("error")
	}

	result := resp.String()
	if errs := json.Unmarshal([]byte(result), resUser); errs != nil {
		return nil, errs
	}

	return &models.ResponseUser{
		Nama:      resUser.Name,
		Pekerjaan: resUser.Job,
		ID:        resUser.ID,
		CreatedAt: resUser.CreatedAt}, nil
}
