package placeorder

import (
	"context"
	"fmt"

	checkoutDomain "github.com/MarcoBuarque/monolito/internal/modules/checkout/domain"
	"github.com/MarcoBuarque/monolito/internal/modules/checkout/repository"
	clientFacade "github.com/MarcoBuarque/monolito/internal/modules/client_adm/facade"
	productFacade "github.com/MarcoBuarque/monolito/internal/modules/product_adm/facade"
	valueobject "github.com/MarcoBuarque/monolito/internal/modules/shared/domain/value_object"
	catalogFacade "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/facade"
)

type PlaceOrderUseCase struct {
	repo repository.ICheckoutRepository

	clientFacade  clientFacade.IClientAdmFacade
	catalogFacade catalogFacade.IStoreCatalogFacade
	productFacade productFacade.IProductAdmFacade
}

func NewPlaceOrderUseCase(
	repo repository.ICheckoutRepository,
	clientFacade clientFacade.IClientAdmFacade,
	catalogFacade catalogFacade.IStoreCatalogFacade,
	productFacade productFacade.IProductAdmFacade,
) PlaceOrderUseCase {
	return PlaceOrderUseCase{repo, clientFacade, catalogFacade, productFacade}
}

func (controller PlaceOrderUseCase) Execute(ctx context.Context, clientID string, products []ProductInfo) (repository.Order, error) {
	client, err := controller.clientFacade.GetClient(ctx, clientID)
	if err != nil {
		// TODO: ADD LOGS
		return repository.Order{}, err
	}

	addr, err := valueobject.NewAddress(client.Street, client.Number, client.Complement, client.City, client.State, client.ZipCode)
	if err != nil {
		// TODO: ADD LOGS
		return repository.Order{}, err
	}

	clientEntity, err := checkoutDomain.NewClient(client.ID, client.Name, client.Email, addr)
	if err != nil {
		// TODO: ADD LOGS
		return repository.Order{}, err
	}

	if err := controller.validateProducts(ctx, products); err != nil {
		// TODO: ADD LOGS
		return repository.Order{}, err
	}

	productsEntity, err := controller.getProducts(ctx, products)
	if err != nil {
		// TODO: ADD LOGS
		return repository.Order{}, err
	}

	order := checkoutDomain.NewOrderEntity("", "", clientEntity, productsEntity)

	orderData := repository.Convert(order)

	if err := controller.repo.CreateOrder(ctx, orderData); err != nil {
		return repository.Order{}, err
	}

	// TODO: create invoice (QUANDO A INVOICE É PAGA O STATUS DA ORDEM MUDA PARA DONE)
	return orderData, nil

}

func (controller PlaceOrderUseCase) validateProducts(ctx context.Context, products []ProductInfo) error {
	if len(products) == 0 {
		return fmt.Errorf("empty slice of products")
	}

	for _, product := range products {
		stock, err := controller.productFacade.CheckStock(ctx, product.ProductID)
		if err != nil {
			return err
		}

		if stock < product.Quantity {
			return fmt.Errorf("%s product quantity %d is greater than stock %d", product.ProductID, product.Quantity, stock)
		}
	}

	return nil
}
func (controller PlaceOrderUseCase) getProducts(ctx context.Context, products []ProductInfo) ([]checkoutDomain.ProductEntity, error) {
	productsEntity := make([]checkoutDomain.ProductEntity, len(products))
	for i := 0; i < len(products); i++ {
		productData, err := controller.catalogFacade.GetProduct(ctx, products[i].ProductID)
		if err != nil {
			return []checkoutDomain.ProductEntity{}, err
		}

		productsEntity[i], err = checkoutDomain.NewProduct(productData.ID, productData.Name, productData.Description, productData.SalesPrice, products[i].Quantity)
		if err != nil {
			return []checkoutDomain.ProductEntity{}, err
		}
	}

	return productsEntity, nil
}
