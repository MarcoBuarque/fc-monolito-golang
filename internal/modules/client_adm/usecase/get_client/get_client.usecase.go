package getclient

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/repository"
)

// Controller
type GetClientUseCase struct {
	productRepository repository.IClientRepository
}

func NewGetClientUseCase(repository repository.IClientRepository) GetClientUseCase {
	return GetClientUseCase{productRepository: repository}
}

func (controller GetClientUseCase) Execute(ctx context.Context, clientID string) (repository.Client, error) {
	response, err := controller.productRepository.GetClient(ctx, clientID)
	if err != nil {
		// TODO: add log
		return repository.Client{}, err
	}

	return response, nil
}
