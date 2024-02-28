package listproducts

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
)

// Controller
type ListProductsUseCase struct {
	productRepository repository.IProductRepository
}

func NewListProductsUseCase(repository repository.IProductRepository) ListProductsUseCase {
	return ListProductsUseCase{productRepository: repository}
}

func (controller ListProductsUseCase) Execute(ctx context.Context) ([]repository.Product, error) {
	response, err := controller.productRepository.ListProducts(ctx)
	if err != nil {
		// TODO: add log
		return []repository.Product{}, err
	}

	return response, nil
}
