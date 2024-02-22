package findproduct

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
)

// Controller
type FindProductsUseCase struct {
	productRepository repository.IProductRepository
}

func NewFindProductUseCase(repository repository.IProductRepository) FindProductsUseCase {
	return FindProductsUseCase{productRepository: repository}
}

func (controller FindProductsUseCase) Execute(ctx context.Context, productID string) (repository.ProductData, error) {
	response, err := controller.productRepository.Find(ctx, productID)
	if err != nil {
		// TODO: add log
		return repository.ProductData{}, err
	}

	return response, nil
}
