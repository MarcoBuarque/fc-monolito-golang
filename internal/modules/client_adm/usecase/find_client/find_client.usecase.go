package findclient

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
)

// Controller
type FindClientUseCase struct {
	productRepository repository.IClientRepository
}

func NewFindClientUseCase(repository repository.IClientRepository) FindClientUseCase {
	return FindClientUseCase{productRepository: repository}
}

func (controller FindClientUseCase) Execute(ctx context.Context, clientID string) (repository.Client, error) {
	response, err := controller.productRepository.Find(ctx, clientID)
	if err != nil {
		// TODO: add log
		return repository.Client{}, err
	}

	return response, nil
}
