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
	facade             IStoreCatalogFacade
	findAllUseCaseMock = &usecasemocks.IFindAllProductsUseCase{}
	findUseCaseMock    = &usecasemocks.IFindProductUseCase{}
)

func TestMain(m *testing.M) {
	facade = StoreCatalogFacade{findAllUseCase: findAllUseCaseMock, findUseCase: findUseCaseMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewStoreCatalogFacade(t *testing.T) {
	response := NewStoreCatalogFacade(findAllUseCaseMock, findUseCaseMock)

	assert.Equal(t, facade, response)
}

func TestStoreCatalogFacade_FindAll(t *testing.T) {
	assert := assert.New(t)

	products := []repository.ProductData{
		{ID: "xpto_id",
			Name:        "xpto",
			Description: "xpto_description"},
	}

	type expect struct {
		data []repository.ProductData
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
				findAllUseCaseMock.On("Execute", mock.Anything, mock.Anything).Return([]repository.ProductData{}, gorm.ErrInvalidData)
			},
			expect: expect{
				data: []repository.ProductData{},
				err:  gorm.ErrInvalidData,
			},
		},
		{
			title: "Success",
			setupMock: func() {
				// clean mock queue
				findAllUseCaseMock.Mock = mock.Mock{}

				findAllUseCaseMock.On("Execute", mock.Anything, mock.Anything).Return(products, nil)
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

			response, err := facade.FindAll(context.TODO())
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

func TestStoreCatalogFacade_Find(t *testing.T) {
	assert := assert.New(t)

	product := repository.ProductData{
		ID:          "xpto_id",
		Name:        "xpto",
		Description: "xpto_description",
	}

	type args struct {
		ctx       context.Context
		productID string
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
				ctx:       context.Background(),
				productID: "product",
			},
			setupMock: func() {
				findUseCaseMock.On("Execute", mock.Anything, mock.Anything).Return(repository.ProductData{}, gorm.ErrInvalidData)
			},
			expect: expect{
				data: repository.ProductData{},
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

			response, err := facade.Find(tt.args.ctx, tt.args.productID)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.ID, response.ID)
			assert.Equal(tt.expect.data.Name, response.Name)
			assert.Equal(tt.expect.data.SalesPrice, response.SalesPrice)
		})
	}
}
