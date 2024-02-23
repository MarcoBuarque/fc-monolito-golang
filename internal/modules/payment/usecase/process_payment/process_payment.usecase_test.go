package processpayment

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/MarcoBuarque/monolito/internal/modules/payment/mocks/repomocks"
	"github.com/MarcoBuarque/monolito/internal/modules/payment/repository"
	"github.com/MarcoBuarque/monolito/internal/modules/shared/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	useCase  ProcessPaymentUseCase
	repoMock = &repomocks.ITransactionRepository{}
)

func TestMain(m *testing.M) {
	useCase = ProcessPaymentUseCase{transactionRepository: repoMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewProcessPaymentUseCase(t *testing.T) {
	response := NewProcessPaymentUseCase(repoMock)

	assert.Equal(t, useCase, response)
}

func TestProcessPaymentUseCase_Execute(t *testing.T) {
	assert := assert.New(t)

	tx := repository.Transaction{
		OrderID: "product",
		Status:  types.Pending,
		Amount:  decimal.NewFromInt(100),
	}

	type args struct {
		ctx     context.Context
		orderID string
		amount  decimal.Decimal
	}

	type expect struct {
		data repository.Transaction
		err  error
	}

	tests := []struct {
		title     string
		args      args
		setupMock func()
		expect    expect
	}{
		{
			title: "should return an error for an invalid product",
			args: args{
				ctx:     context.Background(),
				orderID: "",
			},
			setupMock: func() {},
			expect: expect{
				data: repository.Transaction{},
				err:  fmt.Errorf("transactionEntity: orderID cannot be empty"),
			},
		},
		{
			title: "should return an error from repository",
			args: args{
				ctx:     context.Background(),
				orderID: "product",
				amount:  tx.Amount,
			},
			setupMock: func() {
				repoMock.On("Save", mock.Anything, mock.Anything).Return(gorm.ErrInvalidData)
			},
			expect: expect{
				data: repository.Transaction{},
				err:  gorm.ErrInvalidData,
			},
		},
		{
			title: "Success",
			args: args{
				ctx:     context.Background(),
				orderID: "product",
				amount:  tx.Amount,
			},
			setupMock: func() {
				// clean mock queue
				repoMock.Mock = mock.Mock{}

				repoMock.On("Save", mock.Anything, mock.Anything).Return(nil)
			},
			expect: expect{
				data: tx,
				err:  nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			response, err := useCase.Execute(tt.args.ctx, tt.args.orderID, tt.args.amount)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.OrderID, response.OrderID)
			assert.Equal(tt.expect.data.Amount, response.Amount)
			assert.Equal(tt.expect.data.Status, response.Status)
		})
	}
}
