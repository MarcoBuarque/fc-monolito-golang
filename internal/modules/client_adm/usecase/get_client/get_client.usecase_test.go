package getclient

import (
	"context"
	"os"
	"testing"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/mocks/repomocks"
	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	useCase  GetClientUseCase
	repoMock = &repomocks.IClientRepository{}
)

func TestMain(m *testing.M) {
	useCase = GetClientUseCase{productRepository: repoMock}
	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewFindProductUseCase(t *testing.T) {
	response := NewGetClientUseCase(repoMock)

	assert.Equal(t, useCase, response)
}

func TestFindProductUseCase_Execute(t *testing.T) {
	assert := assert.New(t)

	product := repository.Client{
		ID:    "xpto_id",
		Name:  "xpto",
		Email: "email",
	}

	type args struct {
		ctx context.Context
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
			title: "should return an error from repository",
			args: args{
				ctx: context.Background(),
			},
			setupMock: func() {
				repoMock.On("Find", mock.Anything, mock.Anything).Return(repository.Client{}, gorm.ErrInvalidData)
			},
			expect: expect{
				data: repository.Client{},
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
			assert.Equal(tt.expect.data.Email, response.Email)
			assert.Equal(tt.expect.data.Street, response.Street)
		})
	}
}
