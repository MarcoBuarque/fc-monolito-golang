package repository

import (
	"time"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ID            string          `gorm:"primarykey" json:"id"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	PurchasePrice decimal.Decimal `json:"purchase_price"`
	SalesPrice    decimal.Decimal `json:"sales_price"`
	Stock         int32           `json:"stock"`

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
