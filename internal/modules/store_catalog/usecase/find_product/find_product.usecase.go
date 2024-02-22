package findproduct

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
)

// Controller
type FindProductUseCase struct {
	productRepository repository.IProductRepository
}

func NewFindProductUseCase(repository repository.IProductRepository) FindProductUseCase {
	return FindProductUseCase{productRepository: repository}
}

func (controller FindProductUseCase) Execute(ctx context.Context, productID string) (repository.Product, error) {
	response, err := controller.productRepository.Find(ctx, productID)
	if err != nil {
		// TODO: add log
		return repository.Product{}, err
	}

	return response, nil
}
