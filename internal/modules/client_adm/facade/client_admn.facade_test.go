package facade

import (
	"context"
	"os"
	"testing"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/mocks/usecasemocks"
	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	facade                  IClientAdmFacade
	createClientUseCaseMock = &usecasemocks.ICreateClientUseCase{}
	getClientUseCaseMock    = &usecasemocks.IGetClientUseCase{}
)

func TestMain(m *testing.M) {
	facade = ClientAdmFacade{createUseCase: createClientUseCaseMock, getUseCase: getClientUseCaseMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewClientAdmFacade(t *testing.T) {
	response := NewClientAdmFacade(createClientUseCaseMock, getClientUseCaseMock)

	assert.Equal(t, facade, response)
}

func TestClientAdmFacade_CreateClient(t *testing.T) {
	assert := assert.New(t)

	product := repository.Client{
		ID:    "xpto_id",
		Name:  "xpto",
		Email: "email",
	}

	type expect struct {
		data repository.Client
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
				createClientUseCaseMock.On("Execute", mock.Anything, mock.Anything).Return(repository.Client{}, gorm.ErrInvalidData)
			},
			expect: expect{
				data: repository.Client{},
				err:  gorm.ErrInvalidData,
			},
		},
		{
			title: "Success",
			setupMock: func() {
				// clean mock queue
				createClientUseCaseMock.Mock = mock.Mock{}

				createClientUseCaseMock.On("Execute", mock.Anything, mock.Anything).Return(product, nil)
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

			response, err := facade.CreateClient(context.TODO(), product)

			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.ID, response.ID)
			assert.Equal(tt.expect.data.Name, response.Name)
			assert.Equal(tt.expect.data.Email, response.Email)
			assert.Equal(tt.expect.data.Street, response.Street)
		})
	}
}

func TestClientAdmFacade_GetClient(t *testing.T) {
	assert := assert.New(t)

	product := repository.Client{
		ID:    "xpto_id",
		Name:  "xpto",
		Email: "xpto_description",
	}

	type args struct {
		ctx       context.Context
		productID string
	}

	type expect struct {
		data repository.Client
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
				getClientUseCaseMock.On("Execute", mock.Anything, mock.Anything).Return(repository.Client{}, gorm.ErrInvalidData)
			},
			expect: expect{
				data: repository.Client{},
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
				getClientUseCaseMock.Mock = mock.Mock{}

				getClientUseCaseMock.On("Execute", mock.Anything, mock.Anything).Return(product, nil)
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

			response, err := facade.GetClient(tt.args.ctx, tt.args.productID)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data.ID, response.ID)
			assert.Equal(tt.expect.data.Name, response.Name)
			assert.Equal(tt.expect.data.Email, response.Email)
			assert.Equal(tt.expect.data.Street, response.Street)
		})
	}
}
