package addclient

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/domain"
	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
)

// Controller
type AddClientUseCase struct {
	productRepository repository.IClientRepository
}

func NewAddClientUseCase(repository repository.IClientRepository) AddClientUseCase {
	return AddClientUseCase{productRepository: repository}
}

func (controller AddClientUseCase) Execute(ctx context.Context, data repository.Client) (repository.Client, error) {
	entity, err := domain.NewClient(data.ID, data.Name, data.Address, data.Email)
	if err != nil {
		// TODO: add log
		return repository.Client{}, err
	}

	if err := controller.productRepository.Add(ctx, entity.ToData()); err != nil {
		// TODO: add log
		return repository.Client{}, err
	}

	return entity.ToData(), nil
}
