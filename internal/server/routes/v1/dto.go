package v1

type UpdateSalesPriceRequest struct {
	Price float32 `binding:"required"`
}
