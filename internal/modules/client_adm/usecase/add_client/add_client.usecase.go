package addclient

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/domain"
	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
	valueobject "github.com/MarcoBuarque/monolito/internal/modules/shared/domain/value_object"
)

// Controller
type AddClientUseCase struct {
	productRepository repository.IClientRepository
}

func NewAddClientUseCase(repository repository.IClientRepository) AddClientUseCase {
	return AddClientUseCase{productRepository: repository}
}

func (controller AddClientUseCase) Execute(ctx context.Context, data repository.Client) (repository.Client, error) {
	address, err := valueobject.NewAddress(data.Street, data.Number, data.Complement, data.City, data.State, data.ZipCode)
	if err != nil {
		// TODO: add log
		return repository.Client{}, err
	}

	entity, err := domain.NewClient(data.ID, data.Name, data.Email, address)
	if err != nil {
		// TODO: add log
		return repository.Client{}, err
	}

	parsedData := repository.Convert(entity)

	if err := controller.productRepository.Add(ctx, parsedData); err != nil {
		// TODO: add log
		return repository.Client{}, err
	}

	return parsedData, nil
}
