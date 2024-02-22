package repository

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProductData struct {
	ID          string `gorm:"primarykey"`
	Name        string
	Description string
	SalesPrice  decimal.Decimal
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (d ProductData) TableName() string { return "products" }
