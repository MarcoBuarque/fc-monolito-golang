package updatesalesprice

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
)

type IUpdateSalesPriceUseCase interface {
	Execute(ctx context.Context, productID string, price float32) (repository.Product, error)
}
