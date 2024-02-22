package repository

import (
	"time"

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
