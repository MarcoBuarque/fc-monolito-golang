package checkstock

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
)

// Controller
type CheckStockUseCase struct {
	productRepository repository.IProductRepository
}

func NewCheckStockUseCase(repository repository.IProductRepository) CheckStockUseCase {
	return CheckStockUseCase{productRepository: repository}
}

func (controller CheckStockUseCase) Execute(ctx context.Context, id string) (int32, error) {

	response, err := controller.productRepository.Find(ctx, id)
	if err != nil {
		// TODO: ADD LOG
		return 0, err
	}

	return response.Stock, nil
}
