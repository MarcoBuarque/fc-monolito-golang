package updatesalesprice

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/mocks/repomocks"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	useCase  UpdateSalesPriceUseCase
	repoMock = &repomocks.IProductRepository{}
)

func TestMain(m *testing.M) {
	useCase = UpdateSalesPriceUseCase{productRepository: repoMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewUpdateSalesPriceUseCase(t *testing.T) {
	response := NewUpdateSalesPriceUseCase(repoMock)

	assert.Equal(t, useCase, response)
}

func TestFindProductUseCase_Execute(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		ctx       context.Context
		productID string
		price     float32
	}

	type expect struct {
		data repository.Product
		err  error
	}

	tests := []struct {
		title     string
		args      args
		setupMock func()
		expect    expect
	}{
		{
			title: "Should return an error when price is invalid",
			args: args{
				ctx: context.Background(),
			},
			setupMock: func() {},
			expect:    expect{repository.Product{}, fmt.Errorf("updateSalesPriceUseCase: price cannot be less than or equal zero")},
		},
		{
			title: "Should return an error from repository when try to get product",
			args: args{
				ctx:   context.Background(),
				price: 10,
			},
			setupMock: func() {
				repoMock.On("GetProduct", mock.Anything, mock.Anything).Return(repository.Product{}, gorm.ErrRecordNotFound)
			},
			expect: expect{repository.Product{}, gorm.ErrRecordNotFound},
		},
		{
			title: "Should return an error from repository when try to update price",
			args: args{
				ctx:   context.Background(),
				price: 10,
			},
			setupMock: func() {
				repoMock.Mock = mock.Mock{}
				repoMock.On("GetProduct", mock.Anything, mock.Anything).Return(repository.Product{}, nil)
				repoMock.On("UpdateSalesPrice", mock.Anything, mock.Anything, mock.Anything).Return(gorm.ErrInvalidData)
			},
			expect: expect{repository.Product{}, gorm.ErrInvalidData},
		},
		{
			title: "Success",
			args: args{
				ctx:   context.Background(),
				price: 10,
			},
			setupMock: func() {
				// clean mock queue
				repoMock.Mock = mock.Mock{}

				repoMock.On("GetProduct", mock.Anything, mock.Anything).Return(repository.Product{}, nil)
				repoMock.On("UpdateSalesPrice", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
			expect: expect{repository.Product{SalesPrice: decimal.NewFromInt(10)}, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			response, err := useCase.Execute(tt.args.ctx, tt.args.productID, tt.args.price)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.ID, response.ID)
			assert.Equal(tt.expect.data.Name, response.Name)
			assert.Equal(tt.expect.data.SalesPrice.String(), response.SalesPrice.String())

		})
	}
}
