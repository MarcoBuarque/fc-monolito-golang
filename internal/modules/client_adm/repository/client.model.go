package repository

import (
	"time"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/domain"
	"github.com/MarcoBuarque/monolito/internal/modules/shared/types"
	"gorm.io/gorm"
)

type Client struct {
	ID    string `gorm:"primarykey"`
	Name  string `binding:"required"`
	Email string `binding:"required"`

	// Document
	DocumentType   types.DocumentType `binding:"required"`
	DocumentNumber string             `binding:"required"`

	// Address
	Street     string `binding:"required"`
	Number     string `binding:"required"`
	Complement string `binding:"required"`
	City       string `binding:"required"`
	State      string `binding:"required"`
	ZipCode    string `binding:"required"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func Convert(data domain.ClientEntity) Client {
	return Client{
		ID:    string(data.ID().ToString()),
		Name:  data.Name(),
		Email: data.Email(),
		// Document
		DocumentNumber: data.Document().Number(),
		DocumentType:   data.Document().DocumentType(),

		// Address
		Street:     data.Address().Street(),
		Number:     data.Address().Number(),
		Complement: data.Address().Complement(),
		City:       data.Address().City(),
		State:      data.Address().State(),
		ZipCode:    data.Address().ZipCode(),

		CreatedAt: data.CreatedAt(),
		UpdatedAt: data.UpdatedAt(),
	}
}
