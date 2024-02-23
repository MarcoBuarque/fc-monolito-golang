package updatesalesprice

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/mocks/repomocks"
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

func TestNewFindProductUseCase(t *testing.T) {
	response := NewUpdateSalesPriceUseCase(repoMock)

	assert.Equal(t, useCase, response)
}

func TestFindProductUseCase_Execute(t *testing.T) {
	assert := assert.New(t)

	type args struct {
		ctx       context.Context
		productID string
		price     float64
	}

	tests := []struct {
		title     string
		args      args
		setupMock func()
		expect    error
	}{
		{
			title: "should return an error when price is invalid",
			args: args{
				ctx: context.Background(),
			},
			setupMock: func() {},
			expect:    fmt.Errorf("updateSalesPriceUseCase: price cannot be less than or equal zero"),
		},
		{
			title: "should return an error from repository",
			args: args{
				ctx:   context.Background(),
				price: 10,
			},
			setupMock: func() {
				repoMock.On("UpdateSalesPrice", mock.Anything, mock.Anything, mock.Anything).Return(gorm.ErrInvalidData)
			},
			expect: gorm.ErrInvalidData,
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

				repoMock.On("UpdateSalesPrice", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
			expect: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			err := useCase.Execute(tt.args.ctx, tt.args.productID, tt.args.price)
			assert.Equal(tt.expect, err)

		})
	}
}
