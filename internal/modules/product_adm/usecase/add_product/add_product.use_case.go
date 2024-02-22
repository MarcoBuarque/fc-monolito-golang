package addproduct

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/domain"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
)

// Controller
type AddProductUseCase struct {
	productRepository repository.IProductRepository
}

func NewAddProductUseCase(repository repository.IProductRepository) AddProductUseCase {
	return AddProductUseCase{productRepository: repository}
}

func (controller AddProductUseCase) Execute(ctx context.Context, data repository.Product) (repository.Product, error) {
	entity, err := domain.NewProduct(data.ID, data.Name, data.Description, data.PurchasePrice, data.Stock)
	if err != nil {
		// TODO: add log
		return repository.Product{}, err
	}

	convertedData := repository.Convert(entity)

	if err := controller.productRepository.Add(ctx, convertedData); err != nil {
		// TODO: add log
		return repository.Product{}, err
	}

	return convertedData, nil
}
