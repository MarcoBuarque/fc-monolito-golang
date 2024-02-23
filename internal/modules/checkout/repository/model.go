package repository

import (
	"github.com/MarcoBuarque/monolito/internal/modules/checkout/domain"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID          string `gorm:"primarykey"`
	Name        string
	Description string
	SalesPrice  decimal.Decimal
}

type Client struct {
	ID    string `gorm:"primarykey"`
	Name  string
	Email string

	// Address
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

		// Address
		Street:     data.Address().Street(),
		Number:     data.Address().Number(),
		Complement: data.Address().Complement(),
		City:       data.Address().City(),
		State:      data.Address().State(),
		ZipCode:    data.Address().ZipCode(),
	}
}
