package facade

import (
	"context"
	"os"
	"testing"

	"github.com/MarcoBuarque/monolito/internal/mocks/usecasemocks"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	facade      ProductAdmFacade
	useCaseMock = &usecasemocks.IAddProductUseCase{}
)

func TestMain(m *testing.M) {
	facade = ProductAdmFacade{addUseCase: useCaseMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewProductAdmFacade(t *testing.T) {
	response := NewProductAdmFacade(useCaseMock)

	assert.Equal(t, facade, response)
}

func TestProductAdmFacade_AddProduct(t *testing.T) {
	assert := assert.New(t)

	product := repository.ProductData{
		ID:            "xpto_id",
		Name:          "xpto",
		Description:   "xpto_description",
		PurchasePrice: decimal.NewFromInt(10),
		Stock:         10,
	}

	type args struct {
		ctx     context.Context
		product repository.ProductData
	}

	type expect struct {
		data repository.ProductData
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
				product: product,
			},
			setupMock: func() {
				useCaseMock.On("Execute", mock.Anything, mock.Anything).Return(repository.ProductData{}, gorm.ErrInvalidData)
			},
			expect: expect{
				data: repository.ProductData{},
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
				useCaseMock.Mock = mock.Mock{}

				useCaseMock.On("Execute", mock.Anything, mock.Anything).Return(product, nil)
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

			response, err := facade.AddProduct(tt.args.ctx, tt.args.product)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.ID, response.ID)
			assert.Equal(tt.expect.data.Name, response.Name)
			assert.Equal(tt.expect.data.PurchasePrice, response.PurchasePrice)
			assert.Equal(tt.expect.data.Stock, response.Stock)
		})
	}
}
