package checkstock

import (
	"context"
	"os"
	"testing"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/mocks/repomocks"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	useCase  CheckStockUseCase
	repoMock = &repomocks.IProductRepository{}
)

func TestMain(m *testing.M) {
	useCase = CheckStockUseCase{productRepository: repoMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewCheckStockUseCase(t *testing.T) {
	response := NewCheckStockUseCase(repoMock)

	assert.Equal(t, useCase, response)
}

func TestCheckStockUseCase_Execute(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		ctx       context.Context
		productID string
	}

	type expect struct {
		stock int32
		err   error
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
				ctx:       context.Background(),
				productID: "",
			},
			setupMock: func() {
				repoMock.On("GetProduct", mock.Anything, mock.Anything).Return(repository.Product{}, gorm.ErrInvalidData)
			},
			expect: expect{
				stock: 0,
				err:   gorm.ErrInvalidData,
			},
		},
		{
			title: "Success",
			args: args{
				ctx:       context.Background(),
				productID: "",
			},
			setupMock: func() {
				// clean mock queue
				repoMock.Mock = mock.Mock{}

				repoMock.On("GetProduct", mock.Anything, mock.Anything).Return(repository.Product{Stock: 2}, nil)
			},
			expect: expect{
				stock: 2,
				err:   nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			response, err := useCase.Execute(tt.args.ctx, tt.args.productID)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.stock, response)
		})
	}
}
