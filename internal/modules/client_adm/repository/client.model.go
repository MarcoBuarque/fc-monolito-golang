package repository

import (
	"time"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/domain"
	"gorm.io/gorm"
)

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
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func Convert(data domain.ClientEntity) Client {
	return Client{
		ID:        string(data.ID().ToString()),
		Name:      data.Name(),
		Email:     data.Email(),
		CreatedAt: data.CreatedAt(),
		UpdatedAt: data.UpdatedAt(),
	}
}
