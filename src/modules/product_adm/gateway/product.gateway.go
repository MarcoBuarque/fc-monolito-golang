package gateway

import (
	"context"

	"github.com/MarcoBuarque/monolito/src/modules/product_adm/domain"
)

// Repository/DB Integration
type ProductGateway interface {
	Add(ctx context.Context, data domain.ProductEntity)
	Find(ctx context.Context, id string) domain.ProductEntity
}
