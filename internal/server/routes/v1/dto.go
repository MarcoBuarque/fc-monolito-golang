package v1

import (
	placeorder "github.com/MarcoBuarque/monolito/internal/modules/checkout/usecase/place_order"
)

type UpdateSalesPriceRequest struct {
	Price float32 `binding:"required"`
}

type CheckoutRequest struct {
	ClientID string `binding:"required"`
	Products []placeorder.ProductInfo
}
