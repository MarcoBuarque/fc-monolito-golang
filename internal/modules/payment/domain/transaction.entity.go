package domain

import (
	"fmt"

	"github.com/MarcoBuarque/monolito/internal/modules/payment/repository"
	"github.com/MarcoBuarque/monolito/internal/modules/shared/domain/entity"
	valueobject "github.com/MarcoBuarque/monolito/internal/modules/shared/domain/value_object"
	"github.com/shopspring/decimal"
)

type TransactionEntity struct {
	orderID string
	amount  decimal.Decimal
	status  repository.Status
	entity.BaseEntity
}

func NewTransaction(id, orderID, status string, amount decimal.Decimal) (TransactionEntity, error) {
	if orderID == "" {
		return TransactionEntity{}, fmt.Errorf("transactionEntity: orderID cannot be empty")
	}

	if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		return TransactionEntity{}, fmt.Errorf("transactionEntity: amount cannot be less than or equal zero")
	}

	parsedStatus := repository.Pending
	if status != "" {
		parsedStatus = repository.Status(status)
	}

	return TransactionEntity{
		orderID:    orderID,
		amount:     amount,
		status:     parsedStatus,
		BaseEntity: entity.CreateBaseEntity(valueobject.CreateID(id)),
	}, nil
}

func (data TransactionEntity) Status() string          { return data.status.ToString() }
func (data TransactionEntity) OrderID() string         { return data.orderID }
func (data TransactionEntity) Amount() decimal.Decimal { return data.amount }

func (data *TransactionEntity) Approve() { data.status = repository.Approved }
func (data *TransactionEntity) Decline() { data.status = repository.Declined }

func (data *TransactionEntity) Process() {
	if data.amount.LessThan(decimal.NewFromInt(100)) {
		data.Decline()
		return
	}

	data.Approve()
}

func (data TransactionEntity) ToData() repository.Transaction {
	return repository.Transaction{
		ID:        string(data.ID().ToString()),
		Status:    data.status,
		OrderID:   data.orderID,
		Amount:    data.amount,
		CreatedAt: data.CreatedAt(),
		UpdatedAt: data.UpdatedAt(),
	}
}
