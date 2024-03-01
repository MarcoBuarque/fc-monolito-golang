package facade

import (
	"context"
	"os"
	"testing"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/payment/mocks/usecasemocks"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/payment/repository"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	facade         PaymentFacade
	processPayMock = &usecasemocks.IProcessPaymentUseCase{}
)

func TestMain(m *testing.M) {
	facade = PaymentFacade{processPayMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewPaymentFacade(t *testing.T) {
	response := NewPaymentFacade(processPayMock)

	assert.Equal(t, facade, response)
}

func TestPaymentFacade_Process(t *testing.T) {
	assert := assert.New(t)

	product := repository.Transaction{
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
			title: "Should return an error from repository",
			args: args{
				ctx:     context.Background(),
				orderID: "product",
			},
			setupMock: func() {
				processPayMock.On("Execute", mock.Anything, mock.Anything, mock.Anything).Return(repository.Transaction{}, gorm.ErrInvalidData)
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
			},
			setupMock: func() {
				// clean mock queue
				processPayMock.Mock = mock.Mock{}

				processPayMock.On("Execute", mock.Anything, mock.Anything, mock.Anything).Return(product, nil)
			},
			expect: expect{
				data: product,
				err:  nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			response, err := facade.Process(tt.args.ctx, tt.args.orderID, tt.args.amount)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.ID, response.ID)
			assert.Equal(tt.expect.data.OrderID, response.OrderID)
			assert.Equal(tt.expect.data.Amount, response.Amount)
			assert.Equal(tt.expect.data.Status, response.Status)
		})
	}
}
