package addproduct

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
	useCase  AddProductUseCase
	repoMock = &repomocks.IProductRepository{}
)

func TestMain(m *testing.M) {
	useCase = AddProductUseCase{productRepository: repoMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewAddProductUseCase(t *testing.T) {
	response := NewAddProductUseCase(repoMock)

	assert.Equal(t, useCase, response)
}

func TestAddProductUseCase_Execute(t *testing.T) {
	assert := assert.New(t)

	product := repository.Product{
		ID:            "xpto_id",
		Name:          "xpto",
		Description:   "xpto_description",
		PurchasePrice: decimal.NewFromInt(10),
		Stock:         10,
	}

	type args struct {
		ctx     context.Context
		product repository.Product
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
			title: "should return an error for an invalid product",
			args: args{
				ctx:     context.Background(),
				product: repository.Product{},
			},
			setupMock: func() {},
			expect: expect{
				data: repository.Product{},
				err:  fmt.Errorf("productEntity: name cannot be empty"),
			},
		},
		{
			title: "should return an error from repository",
			args: args{
				ctx:     context.Background(),
				product: product,
			},
			setupMock: func() {
				repoMock.On("Add", mock.Anything, mock.Anything).Return(gorm.ErrInvalidData)
			},
			expect: expect{
				data: repository.Product{},
				err:  gorm.ErrInvalidData,
			},
		},
		{
			title: "Success",
			args: args{
				ctx:     context.Background(),
				product: product,
			},
			setupMock: func() {
				// clean mock queue
				repoMock.Mock = mock.Mock{}

				repoMock.On("Add", mock.Anything, mock.Anything).Return(nil)
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

			response, err := useCase.Execute(tt.args.ctx, tt.args.product)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.ID, response.ID)
			assert.Equal(tt.expect.data.Name, response.Name)
			assert.Equal(tt.expect.data.PurchasePrice, response.PurchasePrice)
			assert.Equal(tt.expect.data.Stock, response.Stock)
		})
	}
}
