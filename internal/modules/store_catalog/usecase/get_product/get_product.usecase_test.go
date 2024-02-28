package getproduct

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
	useCase  GetProductUseCase
	repoMock = &repomocks.IProductRepository{}
)

func TestMain(m *testing.M) {
	useCase = GetProductUseCase{productRepository: repoMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewGetProductUseCase(t *testing.T) {
	response := NewGetProductUseCase(repoMock)

	assert.Equal(t, useCase, response)
}

func TestGetProductUseCase_Execute(t *testing.T) {
	assert := assert.New(t)

	product := repository.Product{
		ID:          "xpto_id",
		Name:        "xpto",
		Description: "xpto_description",
		SalesPrice:  decimal.NewFromInt(10),
	}

	type args struct {
		ctx context.Context
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
			title: "should return an error from repository",
			args: args{
				ctx: context.Background(),
			},
			setupMock: func() {
				repoMock.On("Find", mock.Anything, mock.Anything).Return(repository.Product{}, gorm.ErrInvalidData)
			},
			expect: expect{
				data: repository.Product{},
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

				repoMock.On("Find", mock.Anything, mock.Anything).Return(product, nil)
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

			response, err := useCase.Execute(tt.args.ctx, "xpto")
			assert.Equal(tt.expect.err, err)

			assert.Equal(tt.expect.data.ID, response.ID)
			assert.Equal(tt.expect.data.Name, response.Name)
			assert.Equal(tt.expect.data.SalesPrice, response.SalesPrice)
		})
	}
}
