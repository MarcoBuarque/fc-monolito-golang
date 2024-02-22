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

func (controller AddProductUseCase) Execute(ctx context.Context, data repository.ProductData) (repository.ProductData, error) {
	entity, err := domain.NewProduct(data.ID, data.Name, data.Description, data.PurchasePrice, data.Stock)
	if err != nil {
		// TODO: add log
		return repository.ProductData{}, err
	}

	if err := controller.productRepository.Add(ctx, entity.ToData()); err != nil {
		// TODO: add log
		return repository.ProductData{}, err
	}

	return entity.ToData(), nil
}
