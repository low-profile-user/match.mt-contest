package models

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Hash      string `json:"hash"`
	Cellphone string `json:"celphone"`
	Status    *bool  `json:"status"`
	Email     string `json:"email"`
}
