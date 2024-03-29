package repository

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProductCatalog struct {
	ID          string `gorm:"primarykey"`
	Name        string
	Description string
	SalesPrice  decimal.Decimal

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (ProductCatalog) TableName() string {
	return "products"
}
