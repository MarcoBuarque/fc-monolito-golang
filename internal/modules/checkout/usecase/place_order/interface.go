package placeorder

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/checkout/repository"
)

type ProductInfo struct {
	Quantity  int32
	ProductID string
}

type IPlaceOrderUseCase interface {
	Execute(ctx context.Context, clientID string, products []ProductInfo) (repository.Order, error)
}
