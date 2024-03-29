package getproduct

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/repository"
)

// Controller
type GetProductUseCase struct {
	productRepository repository.IProductRepository
}

func NewGetProductUseCase(repository repository.IProductRepository) GetProductUseCase {
	return GetProductUseCase{productRepository: repository}
}

func (controller GetProductUseCase) Execute(ctx context.Context, productID string) (repository.ProductCatalog, error) {
	response, err := controller.productRepository.GetProduct(ctx, productID)
	if err != nil {
		// TODO: add log
		return repository.ProductCatalog{}, err
	}

	return response, nil
}
