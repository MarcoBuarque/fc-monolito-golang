package facade

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/repository"
	checkstock "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/usecase/check_stock"
	addproduct "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/usecase/create_product"
	updatesalesprice "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/usecase/update_sales_price"
)

type ProductAdmFacade struct {
	addUseCase         addproduct.ICreateProductUseCase
	checkStockUseCase  checkstock.ICheckStockUseCase
	updatePriceUseCase updatesalesprice.IUpdateSalesPriceUseCase
}

func NewProductAdmFacade(addUseCase addproduct.ICreateProductUseCase, checkStockUseCase checkstock.ICheckStockUseCase, updatePriceUseCase updatesalesprice.IUpdateSalesPriceUseCase) ProductAdmFacade {
	return ProductAdmFacade{addUseCase, checkStockUseCase, updatePriceUseCase}
}

func (facade ProductAdmFacade) CreateProduct(ctx context.Context, data repository.Product) (repository.Product, error) {
	return facade.addUseCase.Execute(ctx, data)
}

func (facade ProductAdmFacade) CheckStock(ctx context.Context, id string) (int32, error) {
	return facade.checkStockUseCase.Execute(ctx, id)
}

func (facade ProductAdmFacade) UpdateSalesPrice(ctx context.Context, productID string, price float32) (repository.Product, error) {
	return facade.updatePriceUseCase.Execute(ctx, productID, price)
}
