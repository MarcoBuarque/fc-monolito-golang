package listproducts

import (
	"context"
	"os"
	"testing"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/mocks/repomocks"
	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	useCase  ListProductsUseCase
	repoMock = &repomocks.IProductRepository{}
)

func TestMain(m *testing.M) {
	useCase = ListProductsUseCase{productRepository: repoMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewListProductsUseCase(t *testing.T) {
	response := NewListProductsUseCase(repoMock)

	assert.Equal(t, useCase, response)
}

func TestListProductsUseCase_Execute(t *testing.T) {
	assert := assert.New(t)

	products := []repository.ProductCatalog{
		{
			ID:          "xpto_id",
			Name:        "xpto",
			Description: "xpto_description",
			SalesPrice:  decimal.NewFromInt(10)},
	}

	type args struct {
		ctx context.Context
	}

	type expect struct {
		data []repository.ProductCatalog
		err  error
	}

	tests := []struct {
		title     string
		args      args
		setupMock func()
		expect    expect
	}{
		{
			title: "should return an error from repository",
			args: args{
				ctx: context.Background(),
			},
			setupMock: func() {
				repoMock.On("FindAll", mock.Anything).Return([]repository.ProductCatalog{}, gorm.ErrInvalidData)
			},
			expect: expect{
				data: []repository.ProductCatalog{},
				err:  gorm.ErrInvalidData,
			},
		},
		{
			title: "Success",
			args: args{
				ctx: context.Background(),
			},
			setupMock: func() {
				// clean mock queue
				repoMock.Mock = mock.Mock{}

				repoMock.On("FindAll", mock.Anything).Return(products, nil)
			},
			expect: expect{
				data: products,
				err:  nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			response, err := useCase.Execute(tt.args.ctx)
			assert.Equal(tt.expect.err, err)
			assert.Equal(len(tt.expect.data), len(response))

			for index, product := range response {
				assert.Equal(tt.expect.data[index].ID, product.ID)
				assert.Equal(tt.expect.data[index].Name, product.Name)
				assert.Equal(tt.expect.data[index].SalesPrice, product.SalesPrice)
			}
		})
	}
}
