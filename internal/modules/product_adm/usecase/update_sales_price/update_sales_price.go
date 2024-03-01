package updatesalesprice

import (
	"context"
	"fmt"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	"github.com/shopspring/decimal"
)

// Controller
type UpdateSalesPriceUseCase struct {
	productRepository repository.IProductRepository
}

func NewUpdateSalesPriceUseCase(repository repository.IProductRepository) UpdateSalesPriceUseCase {
	return UpdateSalesPriceUseCase{productRepository: repository}
}

func (controller UpdateSalesPriceUseCase) Execute(ctx context.Context, productID string, price float32) (repository.Product, error) {
	if price <= 0 {
		return repository.Product{}, fmt.Errorf("updateSalesPriceUseCase: price cannot be less than or equal zero")
	}

	data, err := controller.productRepository.GetProduct(ctx, productID)
	if err != nil {
		// TODO: add log
		return repository.Product{}, err
	}

	data.SalesPrice = decimal.NewFromFloat32(price)

	if err := controller.productRepository.UpdateSalesPrice(ctx, productID, price); err != nil {
		// TODO: add log
		return repository.Product{}, err
	}

	return data, nil
}
