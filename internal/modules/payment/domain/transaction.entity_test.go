package domain

import (
	"fmt"
	"testing"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		id      string
		orderID string
		status  types.Status
		amount  decimal.Decimal
	}

	type expect struct {
		data TransactionEntity
		err  error
	}
	tests := []struct {
		title  string
		args   args
		expect expect
	}{
		{
			title: "Should return an error when orderID is empty",
			args:  args{},
			expect: expect{
				data: TransactionEntity{}, err: fmt.Errorf("transactionEntity: orderID cannot be empty"),
			},
		},

		{
			title: "Should return an error when amount is less than or equal zero",
			args: args{
				orderID: "xpto",
				status:  "xpto_description",
			},
			expect: expect{
				data: TransactionEntity{}, err: fmt.Errorf("transactionEntity: amount cannot be less than or equal zero"),
			},
		},
		{
			title: "Success with pending status",
			args: args{
				id:      "xpto_id",
				orderID: "xpto",
				amount:  decimal.NewFromInt(100),
			},
			expect: expect{
				data: TransactionEntity{
					orderID: "xpto",
					status:  types.Pending,
					amount:  decimal.NewFromInt(100),
				}, err: nil,
			},
		},
		{
			title: "Success without pending status",
			args: args{
				id:      "xpto_id",
				orderID: "xpto",
				status:  "declined",
				amount:  decimal.NewFromInt(70),
			},
			expect: expect{
				data: TransactionEntity{
					orderID: "xpto",
					status:  "declined",
					amount:  decimal.NewFromInt(70),
				}, err: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			response, err := NewTransaction(tt.args.id, tt.args.orderID, tt.args.status, tt.args.amount)

			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.OrderID(), response.OrderID())
			assert.Equal(tt.expect.data.Amount(), response.Amount())
			assert.Equal(tt.expect.data.Status(), response.Status())
		})
	}
}

func TestProcess(t *testing.T) {
	assert := assert.New(t)

	entity := TransactionEntity{
		orderID: "xpto",
		status:  types.Pending,
		amount:  decimal.NewFromInt(70),
	}

	assert.Equal(entity.Status(), types.Pending)
	entity.Process()

	assert.Equal(entity.Status(), types.Declined)

	entity = TransactionEntity{
		orderID: "xpto",
		status:  types.Pending,
		amount:  decimal.NewFromInt(800),
	}
	entity.Process()
	assert.Equal(entity.Status(), types.Approved)
}
