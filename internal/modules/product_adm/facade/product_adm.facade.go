package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	checkstock "github.com/MarcoBuarque/monolito/internal/modules/product_adm/usecase/check_stock"
	addproduct "github.com/MarcoBuarque/monolito/internal/modules/product_adm/usecase/create_product"
)

type ProductAdmFacade struct {
	addUseCase        addproduct.ICreateProductUseCase
	checkStockUseCase checkstock.ICheckStockUseCase
}

func NewProductAdmFacade(addUseCase addproduct.ICreateProductUseCase, checkStockUseCase checkstock.ICheckStockUseCase) ProductAdmFacade {
	return ProductAdmFacade{addUseCase, checkStockUseCase}
}

func (facade ProductAdmFacade) CreateProduct(ctx context.Context, data repository.Product) (repository.Product, error) {
	return facade.addUseCase.Execute(ctx, data)
}

func (facade ProductAdmFacade) CheckStock(ctx context.Context, id string) (int32, error) {
	return facade.checkStockUseCase.Execute(ctx, id)
}
