package updatesalesprice

import (
	"context"
)

type IUpdateSalesPriceUseCase interface {
	Execute(ctx context.Context, productID string, price float64) error
}
