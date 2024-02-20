package addproduct

import (
	"context"

	"github.com/MarcoBuarque/monolito/src/modules/product_adm/domain"
	"github.com/MarcoBuarque/monolito/src/modules/product_adm/gateway"
	"github.com/MarcoBuarque/monolito/src/modules/product_adm/repository"
)

// Controller
type AddProductUseCase struct {
	productRepository gateway.ProductGateway
}

func (controller AddProductUseCase) Execute(ctx context.Context, data repository.ProductData) (repository.ProductData, error) {
	entity, err := domain.NewProduct(data.ID, data.Name, data.Description, data.PurchasePrice, data.Stock)
	if err != nil {
		return repository.ProductData{}, err
	}

	controller.productRepository.Add(ctx, entity)

	return entity.ToData(), nil
}
