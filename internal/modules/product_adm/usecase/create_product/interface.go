package addproduct

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
)

type ICreateProductUseCase interface {
	Execute(ctx context.Context, data repository.Product) (repository.Product, error)
}
