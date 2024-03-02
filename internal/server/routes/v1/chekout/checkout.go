package checkout

import (
	"net/http"

	"github.com/MarcoBuarque/fc-monolito-golang/config"
	"github.com/MarcoBuarque/fc-monolito-golang/constant"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/checkout/repository"
	placeorder "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/checkout/usecase/place_order"
	clientAdmFactory "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/factory"
	productAdmFactory "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/factory"
	productCatalogFactory "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/factory"
	"github.com/MarcoBuarque/fc-monolito-golang/pkg"
	"github.com/gin-gonic/gin"
)

var placeOrderUseCase placeorder.IPlaceOrderUseCase

type checkoutRequest struct {
	ClientID string `binding:"required"`
	Products []placeorder.ProductInfo
}

func init() {
	placeOrderUseCase = placeorder.NewPlaceOrderUseCase(
		repository.NewCheckoutRepository(config.GetDB()),
		clientAdmFactory.NewClientAdmFacadeFactory(),
		productCatalogFactory.NewStoreCatalogFacadeFactory(),
		productAdmFactory.NewProductAdmFacadeFactory(),
	)
}

func ConfigRoutes(v1 *gin.RouterGroup) {
	checkoutGroup := v1.Group("/checkout")
	{
		checkoutGroup.POST("", checkout)
	}
}

// Checkout godoc
//
//	@Summary		Create a new order.
//	@Description	This endpoint allows you to create a new order by providing its details in the request body.
//	@Tags			checkout
//	@Accept			json
//	@Produce		json
//	@Param			order	body		checkoutRequest	true	"Checkout data"
//	@Success		200		{object}	pkg.ApiResponse[repository.Order]
//
//	@Failure		400		{object}	pkg.ApiResponse[pkg.Null]
//	@Failure		500		{object}	pkg.ApiResponse[pkg.Null]
//
//	@Router			/checkout [post]
func checkout(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request checkoutRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	response, err := placeOrderUseCase.Execute(c.Request.Context(), request.ClientID, request.Products)

	if err != nil {
		pkg.GormErrorHandler("Routes V1: Checkout:", err)
	}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, response))
}
