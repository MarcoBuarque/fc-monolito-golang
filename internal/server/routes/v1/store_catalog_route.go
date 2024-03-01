package v1

import (
	"net/http"

	"github.com/MarcoBuarque/monolito/constant"

	productCatalogFactory "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/factory"
	repoProductCatalog "github.com/MarcoBuarque/monolito/internal/modules/store_catalog/repository"
	"github.com/MarcoBuarque/monolito/pkg"
	"github.com/gin-gonic/gin"
)

var _ repoProductCatalog.ProductCatalog

// GetProduct godoc
//
//	@Summary		Retrieve a specific product by its ID.
//	@Description	This endpoint retrieves the details of a single product based on the provided product ID.
//	@Tags			catalogs
//	@Accept			json
//	@Produce		json
//
//	@Param			productID	path		string	true	"Product ID"
//
//	@Success		200			{object}	pkg.ApiResponse[repoProductCatalog.ProductCatalog]
//
//	@Failure		400			{object}	pkg.ApiResponse[pkg.Null]
//	@Failure		500			{object}	pkg.ApiResponse[pkg.Null]
//
//	@Router			/store-catalog/products/{productID} [get]
func GetProduct(c *gin.Context) {
	defer pkg.PanicHandler(c)

	productID := c.Param("productID")

	response, err := productCatalogFactory.NewStoreCatalogFacadeFactory().GetProduct(c.Request.Context(), productID)
	if err != nil {
		pkg.GormErrorHandler("Routes V1: GetProduct", err)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

// ListProducts godoc
//
//	@Summary		Retrieve a list of active products.
//	@Description	This endpoint returns a list of products with information such as name, price, description, and status.
//	@Tags			catalogs
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	pkg.ApiResponse[[]repoProductCatalog.ProductCatalog]
//
//	@Failure		400	{object}	pkg.ApiResponse[pkg.Null]
//	@Failure		500	{object}	pkg.ApiResponse[pkg.Null]
//
//	@Router			/store-catalog/products/ [get]
func ListProducts(c *gin.Context) {
	defer pkg.PanicHandler(c)

	response, err := productCatalogFactory.NewStoreCatalogFacadeFactory().ListProducts(c.Request.Context())
	if err != nil {
		pkg.GormErrorHandler("Routes V1: ListProducts", err)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}
