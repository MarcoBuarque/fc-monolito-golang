package addproduct

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
)

type IAddProductUseCase interface {
	Execute(ctx context.Context, data repository.Product) (repository.Product, error)
}
