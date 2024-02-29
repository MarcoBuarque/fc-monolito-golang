package addproduct

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/domain"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
)

// Controller
type CreateProductUseCase struct {
	productRepository repository.IProductRepository
}

func NewCreateProductUseCase(repository repository.IProductRepository) CreateProductUseCase {
	return CreateProductUseCase{productRepository: repository}
}

func (controller CreateProductUseCase) Execute(ctx context.Context, data repository.Product) (repository.Product, error) {
	entity, err := domain.NewProduct(data.ID, data.Name, data.Description, data.PurchasePrice, data.Stock)
	if err != nil {
		// TODO: add log
		return repository.Product{}, err
	}

	convertedData := repository.Convert(entity)

	if err := controller.productRepository.CreateProduct(ctx, convertedData); err != nil {
		// TODO: add log
		return repository.Product{}, err
	}

	return convertedData, nil
}
