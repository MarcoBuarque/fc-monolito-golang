package repository

import (
	"time"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID            string `gorm:"primarykey"`
	Name          string
	Description   string
	PurchasePrice decimal.Decimal
	Stock         int32
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

func Convert(data domain.ProductEntity) Product {
	return Product{
		ID:            string(data.ID().ToString()),
		Name:          data.Name(),
		Description:   data.Description(),
		PurchasePrice: data.PurchasePrice(),
		Stock:         data.Stock(),
		CreatedAt:     data.CreatedAt(),
		UpdatedAt:     data.UpdatedAt(),
	}
}
