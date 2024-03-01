package repository

import (
	"time"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID            string          `gorm:"primarykey" json:"id" binding:"required"`
	Name          string          `json:"name" binding:"required"`
	Description   string          `json:"description" binding:"required"`
	PurchasePrice decimal.Decimal `json:"purchase_price" binding:"required"`
	SalesPrice    decimal.Decimal `json:"sales_price" binding:"required"`
	Stock         int32           `json:"stock" binding:"required"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func Convert(data domain.ProductEntity) Product {
	return Product{
		ID:            string(data.ID().ToString()),
		Name:          data.Name(),
		Description:   data.Description(),
		PurchasePrice: data.PurchasePrice(),
		SalesPrice:    data.PurchasePrice(),
		Stock:         data.Stock(),
		CreatedAt:     data.CreatedAt(),
		UpdatedAt:     data.UpdatedAt(),
	}
}
