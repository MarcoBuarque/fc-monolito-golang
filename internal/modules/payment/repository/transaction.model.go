package repository

import (
	"time"

	"github.com/MarcoBuarque/monolito/internal/modules/payment/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Declined Status = "declined"
)

type Transaction struct {
	ID        string `gorm:"primarykey"`
	OrderID   string
	Amount    decimal.Decimal
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (data Status) ToString() string {
	return string(data)
}

func Convert(data domain.TransactionEntity) Transaction {
	return Transaction{
		ID:        string(data.ID().ToString()),
		Status:    Status(data.Status()),
		OrderID:   data.OrderID(),
		Amount:    data.Amount(),
		CreatedAt: data.CreatedAt(),
		UpdatedAt: data.UpdatedAt(),
	}
}
