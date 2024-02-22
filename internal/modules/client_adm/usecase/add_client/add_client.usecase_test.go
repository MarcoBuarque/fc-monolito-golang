package addclient

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/mocks/repomocks"
	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	useCase  AddClientUseCase
	repoMock = &repomocks.IClientRepository{}
)

func TestMain(m *testing.M) {
	useCase = AddClientUseCase{productRepository: repoMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewAddClientUseCase(t *testing.T) {
	response := NewAddClientUseCase(repoMock)

	assert.Equal(t, useCase, response)
}

func TestAddClientUseCase_Execute(t *testing.T) {
	assert := assert.New(t)

	product := repository.Client{
		ID:    "xpto_id",
		Name:  "xpto",
		Email: "email",
	}

	type args struct {
		ctx     context.Context
		product repository.Client
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
			title: "should return an error for an invalid product",
			args: args{
				ctx:     context.Background(),
				product: repository.Client{},
			},
			setupMock: func() {},
			expect: expect{
				data: repository.Client{},
				err:  fmt.Errorf("clientEntity: name cannot be empty"),
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
				data: repository.Client{},
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
			assert.Equal(tt.expect.data.Email, response.Email)
		})
	}
}
