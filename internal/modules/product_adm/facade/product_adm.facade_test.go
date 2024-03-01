package facade

import (
	"context"
	"os"
	"testing"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/mocks/usecasemocks"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/repository"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	facade                 ProductAdmFacade
	addProductMock         = &usecasemocks.ICreateProductUseCase{}
	checkStockMock         = &usecasemocks.ICheckStockUseCase{}
	updatePriceUseCaseMock = &usecasemocks.IUpdateSalesPriceUseCase{}
)

func TestMain(m *testing.M) {
	facade = ProductAdmFacade{addUseCase: addProductMock, checkStockUseCase: checkStockMock, updatePriceUseCase: updatePriceUseCaseMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewProductAdmFacade(t *testing.T) {
	response := NewProductAdmFacade(addProductMock, checkStockMock, updatePriceUseCaseMock)

	assert.Equal(t, facade, response)
}

func TestProductAdmFacade_CreateProduct(t *testing.T) {
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
			title: "Should return an error from repository",
			args: args{
				ctx:     context.Background(),
				product: product,
			},
			setupMock: func() {
				addProductMock.On("Execute", mock.Anything, mock.Anything).Return(repository.Product{}, gorm.ErrInvalidData)
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
				addProductMock.Mock = mock.Mock{}

				addProductMock.On("Execute", mock.Anything, mock.Anything).Return(product, nil)
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

			response, err := facade.CreateProduct(tt.args.ctx, tt.args.product)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.ID, response.ID)
			assert.Equal(tt.expect.data.Name, response.Name)
			assert.Equal(tt.expect.data.PurchasePrice, response.PurchasePrice)
			assert.Equal(tt.expect.data.Stock, response.Stock)
		})
	}
}

func TestProductAdmFacade_CheckStock(t *testing.T) {
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
			title: "Should return an error from repository",
			args: args{
				ctx:       context.Background(),
				productID: "product",
			},
			setupMock: func() {
				checkStockMock.On("Execute", mock.Anything, mock.Anything).Return(int32(0), gorm.ErrInvalidData)
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
				productID: "product",
			},
			setupMock: func() {
				// clean mock queue
				checkStockMock.Mock = mock.Mock{}

				checkStockMock.On("Execute", mock.Anything, mock.Anything).Return(int32(10), nil)
			},
			expect: expect{
				stock: 10,
				err:   nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			stock, err := facade.CheckStock(tt.args.ctx, tt.args.productID)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.stock, stock)
		})
	}
}

func TestStoreCatalogFacade_UpdateSalesPrice(t *testing.T) {
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
			title: "Should return an error from repository",
			args: args{
				ctx:       context.Background(),
				productID: "product",
				price:     10,
			},
			setupMock: func() {
				updatePriceUseCaseMock.On("Execute", mock.Anything, mock.Anything, mock.Anything).Return(repository.Product{}, gorm.ErrInvalidData)
			},
			expect: expect{repository.Product{}, gorm.ErrInvalidData},
		},
		{
			title: "Success",
			args: args{
				ctx:       context.Background(),
				productID: "product",
				price:     10,
			},
			setupMock: func() {
				// clean mock queue
				updatePriceUseCaseMock.Mock = mock.Mock{}

				updatePriceUseCaseMock.On("Execute", mock.Anything, mock.Anything, mock.Anything).Return(repository.Product{}, nil)
			},
			expect: expect{repository.Product{}, nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			response, err := facade.UpdateSalesPrice(tt.args.ctx, tt.args.productID, tt.args.price)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.ID, response.ID)
			assert.Equal(tt.expect.data.Name, response.Name)
			assert.Equal(tt.expect.data.SalesPrice.String(), response.SalesPrice.String())

		})
	}
}
