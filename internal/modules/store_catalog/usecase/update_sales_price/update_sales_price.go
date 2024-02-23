package updatesalesprice

import (
	"context"
	"fmt"

	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
)

// Controller
type UpdateSalesPriceUseCase struct {
	productRepository repository.IProductRepository
}

func NewUpdateSalesPriceUseCase(repository repository.IProductRepository) UpdateSalesPriceUseCase {
	return UpdateSalesPriceUseCase{productRepository: repository}
}

func (controller UpdateSalesPriceUseCase) Execute(ctx context.Context, productID string, price float64) error {
	if price <= 0 {
		return fmt.Errorf("updateSalesPriceUseCase: price cannot be less than or equal zero")
	}

	err := controller.productRepository.UpdateSalesPrice(ctx, productID, price)
	if err != nil {
		// TODO: add log
		return err
	}

	return nil
}
