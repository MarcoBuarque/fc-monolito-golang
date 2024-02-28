package repository

import (
	"time"

	"github.com/MarcoBuarque/monolito/internal/modules/payment/domain"
	"github.com/MarcoBuarque/monolito/internal/modules/shared/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Transaction struct {
	ID        string `gorm:"primarykey"`
	OrderID   string
	Amount    decimal.Decimal
	Status    types.Status
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func Convert(data domain.TransactionEntity) Transaction {
	return Transaction{
		ID:        string(data.ID().ToString()),
		Status:    types.Status(data.Status()),
		OrderID:   data.OrderID(),
		Amount:    data.Amount(),
		CreatedAt: data.CreatedAt(),
		UpdatedAt: data.UpdatedAt(),
	}
}
