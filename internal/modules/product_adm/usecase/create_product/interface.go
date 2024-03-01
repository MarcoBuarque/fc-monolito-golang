package addproduct

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/repository"
)

type ICreateProductUseCase interface {
	Execute(ctx context.Context, data repository.Product) (repository.Product, error)
}
