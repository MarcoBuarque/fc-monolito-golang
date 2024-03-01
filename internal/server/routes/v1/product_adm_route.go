package v1

import (
	"net/http"

	"github.com/MarcoBuarque/monolito/constant"
	productAdmFactory "github.com/MarcoBuarque/monolito/internal/modules/product_adm/factory"
	repoProductADM "github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	"github.com/MarcoBuarque/monolito/pkg"
	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
//
//	@Summary		Create a new product.
//	@Description	This endpoint allows you to create a new product by providing its details in the request body.
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		repoProductADM.Product	true	"Product data"
//	@Success		200		{object}	pkg.ApiResponse[repoProductADM.Product]
//
//	@Failure		400		{object}	pkg.ApiResponse[pkg.Null]
//	@Failure		500		{object}	pkg.ApiResponse[pkg.Null]
//
//	@Router			/product-adm [post]
func CreateProduct(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request repoProductADM.Product
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	response, err := productAdmFactory.NewProductAdmFacadeFactory().CreateProduct(c.Request.Context(), request)
	if err != nil {
		pkg.GormErrorHandler("Routes V1: CreateProduct", err)
	}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, response))
}

// UpdateSalesPrice godoc
//
//	@Summary		Update the sales price of a product.
//	@Description	This API allows you to update the sales price of a specific product by providing its ID and the new price in the request body.
//	@Tags			products
//	@Accept			json
//	@Produce		json
//
//	@Param			productID	path	string					true	"Product ID"
//
//	@Param			product		body	UpdateSalesPriceRequest	true	"new price value"

// @Success	200	{object}	pkg.ApiResponse[repoProductADM.Product]
//
// @Failure	400	{object}	pkg.ApiResponse[pkg.Null]
// @Failure	500	{object}	pkg.ApiResponse[pkg.Null]
//
// @Router		/product-adm/{productID} [patch]
func UpdateSalesPrice(c *gin.Context) {
	defer pkg.PanicHandler(c)

	productID := c.Param("productID")

	var request UpdateSalesPriceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	response, err := productAdmFactory.NewProductAdmFacadeFactory().UpdateSalesPrice(c.Request.Context(), productID, request.Price)
	if err != nil {
		pkg.GormErrorHandler("Routes V1: ListProducts", err)
	}

	// TODO: Review: send all product data back?
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}
