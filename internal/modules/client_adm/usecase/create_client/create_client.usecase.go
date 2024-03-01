package createclient

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/domain"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/repository"
	valueobject "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/shared/domain/value_object"
)

// Controller
type CreateClientUseCase struct {
	productRepository repository.IClientRepository
}

func NewCreateClientUseCase(repository repository.IClientRepository) CreateClientUseCase {
	return CreateClientUseCase{productRepository: repository}
}

func (controller CreateClientUseCase) Execute(ctx context.Context, data repository.Client) (repository.Client, error) {
	document, err := valueobject.NewDocument(data.DocumentNumber, data.DocumentType)
	if err != nil {
		// TODO: add log
		return repository.Client{}, err
	}

	address, err := valueobject.NewAddress(data.Street, data.Number, data.Complement, data.City, data.State, data.ZipCode)
	if err != nil {
		// TODO: add log
		return repository.Client{}, err
	}

	entity, err := domain.NewClient(data.ID, data.Name, data.Email, document, address)
	if err != nil {
		// TODO: add log
		return repository.Client{}, err
	}

	parsedData := repository.Convert(entity)

	if err := controller.productRepository.CreateClient(ctx, parsedData); err != nil {
		// TODO: add log
		return repository.Client{}, err
	}

	return parsedData, nil
}
