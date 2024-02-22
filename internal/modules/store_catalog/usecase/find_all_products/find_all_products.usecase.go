package findallproducts

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
)

// Controller
type FindAllProductsUseCase struct {
	productRepository repository.IProductRepository
}

func NewFindAllProductsUseCase(repository repository.IProductRepository) FindAllProductsUseCase {
	return FindAllProductsUseCase{productRepository: repository}
}

func (controller FindAllProductsUseCase) Execute(ctx context.Context) ([]repository.ProductData, error) {
	response, err := controller.productRepository.FindAll(ctx)
	if err != nil {
		// TODO: add log
		return []repository.ProductData{}, err
	}

	return response, nil
}
