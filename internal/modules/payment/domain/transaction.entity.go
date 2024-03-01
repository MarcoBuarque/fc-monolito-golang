package domain

import (
	"fmt"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/domain/entity"
	valueobject "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/domain/value_object"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/types"
	"github.com/shopspring/decimal"
)

type TransactionEntity struct {
	orderID string
	amount  decimal.Decimal
	status  types.Status
	entity.BaseEntity
}

func NewTransaction(id, orderID string, status types.Status, amount decimal.Decimal) (TransactionEntity, error) {
	if orderID == "" {
		return TransactionEntity{}, fmt.Errorf("transactionEntity: orderID cannot be empty")
	}

	if amount.LessThanOrEqual(decimal.NewFromInt(0)) {
		return TransactionEntity{}, fmt.Errorf("transactionEntity: amount cannot be less than or equal zero")
	}

	if status == "" {
		status = types.Pending
	}

	return TransactionEntity{
		orderID:    orderID,
		amount:     amount,
		status:     status,
		BaseEntity: entity.NewBaseEntity(valueobject.NewID(id)),
	}, nil
}

func (data TransactionEntity) Status() types.Status    { return data.status }
func (data TransactionEntity) OrderID() string         { return data.orderID }
func (data TransactionEntity) Amount() decimal.Decimal { return data.amount }

func (data *TransactionEntity) Approve() { data.status = "approved" }
func (data *TransactionEntity) Decline() { data.status = "declined" }

func (data *TransactionEntity) Process() {
	if data.amount.LessThan(decimal.NewFromInt(100)) {
		data.Decline()
		return
	}

	data.Approve()
}
