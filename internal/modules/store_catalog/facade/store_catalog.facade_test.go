package facade

import (
	"context"
	"os"
	"testing"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/mocks/usecasemocks"
	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	facade                  IStoreCatalogFacade
	listProductsUseCaseMock = &usecasemocks.IListProductsUseCase{}
	findUseCaseMock         = &usecasemocks.IGetProductUseCase{}
	updatePriceUseCaseMock  = &usecasemocks.IUpdateSalesPriceUseCase{}
)

func TestMain(m *testing.M) {
	facade = StoreCatalogFacade{findAllUseCase: listProductsUseCaseMock, findUseCase: findUseCaseMock, updatePriceUseCase: updatePriceUseCaseMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewStoreCatalogFacade(t *testing.T) {
	response := NewStoreCatalogFacade(listProductsUseCaseMock, findUseCaseMock, updatePriceUseCaseMock)

	assert.Equal(t, facade, response)
}

func TestStoreCatalogFacade_ListProducts(t *testing.T) {
	assert := assert.New(t)

	products := []repository.Product{
		{ID: "xpto_id",
			Name:        "xpto",
			Description: "xpto_description"},
	}

	type expect struct {
		data []repository.Product
		err  error
	}

	tests := []struct {
		title     string
		setupMock func()
		expect    expect
	}{
		{
			title: "Should return an error from repository",
			setupMock: func() {
				listProductsUseCaseMock.On("Execute", mock.Anything, mock.Anything).Return([]repository.Product{}, gorm.ErrInvalidData)
			},
			expect: expect{
				data: []repository.Product{},
				err:  gorm.ErrInvalidData,
			},
		},
		{
			title: "Success",
			setupMock: func() {
				// clean mock queue
				listProductsUseCaseMock.Mock = mock.Mock{}

				listProductsUseCaseMock.On("Execute", mock.Anything, mock.Anything).Return(products, nil)
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

			response, err := facade.ListProducts(context.TODO())
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

func TestStoreCatalogFacade_GetProduct(t *testing.T) {
	assert := assert.New(t)

	product := repository.Product{
		ID:          "xpto_id",
		Name:        "xpto",
		Description: "xpto_description",
	}

	type args struct {
		ctx       context.Context
		productID string
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
			},
			setupMock: func() {
				findUseCaseMock.On("Execute", mock.Anything, mock.Anything).Return(repository.Product{}, gorm.ErrInvalidData)
			},
			expect: expect{
				data: repository.Product{},
				err:  gorm.ErrInvalidData,
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
				findUseCaseMock.Mock = mock.Mock{}

				findUseCaseMock.On("Execute", mock.Anything, mock.Anything).Return(product, nil)
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

			response, err := facade.GetProduct(tt.args.ctx, tt.args.productID)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.ID, response.ID)
			assert.Equal(tt.expect.data.Name, response.Name)
			assert.Equal(tt.expect.data.SalesPrice, response.SalesPrice)
		})
	}
}

func TestStoreCatalogFacade_UpdateSalesPrice(t *testing.T) {
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
			title: "Should return an error from repository",
			args: args{
				ctx:       context.Background(),
				productID: "product",
				price:     10,
			},
			setupMock: func() {
				updatePriceUseCaseMock.On("Execute", mock.Anything, mock.Anything, mock.Anything).Return(gorm.ErrInvalidData)
			},
			expect: gorm.ErrInvalidData,
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

				updatePriceUseCaseMock.On("Execute", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			},
			expect: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			err := facade.UpdateSalesPrice(tt.args.ctx, tt.args.productID, tt.args.price)
			assert.Equal(tt.expect, err)

		})
	}
}
