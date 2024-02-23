package repository

import "github.com/MarcoBuarque/monolito/internal/modules/checkout/domain"

type Client struct {
	ID         string `gorm:"primarykey"`
	Name       string
	Email      string
	Street     string
	Number     string
	Complement string
	City       string
	State      string
	ZipCode    string
}

func Convert(data domain.ClientEntity) Client {
	return Client{
		ID:    string(data.ID().ToString()),
		Name:  data.Name(),
		Email: data.Email(),
	}
}
