package models

//your models for request, reponse and table struct define here

type RequestUser struct {
	Nama      string `json:"nama"`
	Pekerjaan string `json:"pekerjaan"`
}

type ResponseUser struct {
	Nama      string `json:"nama"`
	Pekerjaan string `json:"pekerjaan"`
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}

type ResponseUserAPI struct {
	Name      string `json:"name"`
	Job       string `json:"job"`
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
}
