package storecatalog

import (
	"net/http"

	"github.com/MarcoBuarque/fc-monolito-golang/config"
	"github.com/MarcoBuarque/fc-monolito-golang/constant"

	productCatalogFactory "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/factory"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/repository"
	getproduct "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/usecase/get_product"
	listproducts "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/store_catalog/usecase/list_products"
	"github.com/MarcoBuarque/fc-monolito-golang/pkg"
	"github.com/gin-gonic/gin"
)

var (
	_           repository.ProductCatalog
	getUsecase  getproduct.IGetProductUseCase
	listUsecase listproducts.IListProductsUseCase
)

func init() {
	repo := repository.NewProductRepository(config.GetDB())

	listUsecase = listproducts.NewListProductsUseCase(repo)
	getUsecase = getproduct.NewGetProductUseCase(repo)
}

func ConfigRoutes(v1 *gin.RouterGroup) {
	storeCatalog := v1.Group("/store-catalog")
	{
		storeCatalog.GET("/products", listProducts)
		storeCatalog.GET("/products/:productID", getProduct)
	}
}

// getProduct godoc
//
//	@Summary		Retrieve a specific product by its ID.
//	@Description	This endpoint retrieves the details of a single product based on the provided product ID.
//	@Tags			catalogs
//	@Accept			json
//	@Produce		json
//
//	@Param			productID	path		string	true	"Product ID"
//
//	@Success		200			{object}	pkg.ApiResponse[repository.ProductCatalog]
//
//	@Failure		400			{object}	pkg.ApiResponse[pkg.Null]
//	@Failure		500			{object}	pkg.ApiResponse[pkg.Null]
//
//	@Router			/store-catalog/products/{productID} [get]
func getProduct(c *gin.Context) {
	defer pkg.PanicHandler(c)

	productID := c.Param("productID")

	response, err := productCatalogFactory.NewStoreCatalogFacadeFactory().GetProduct(c.Request.Context(), productID)
	if err != nil {
		pkg.GormErrorHandler("Routes V1: GetProduct", err)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

// listProducts godoc
//
//	@Summary		Retrieve a list of active products.
//	@Description	This endpoint returns a list of products with information such as name, price, description, and status.
//	@Tags			catalogs
//	@Accept			json
//	@Produce		json
//
//	@Success		200	{object}	pkg.ApiResponse[[]repository.ProductCatalog]
//
//	@Failure		400	{object}	pkg.ApiResponse[pkg.Null]
//	@Failure		500	{object}	pkg.ApiResponse[pkg.Null]
//
//	@Router			/store-catalog/products/ [get]
func listProducts(c *gin.Context) {
	defer pkg.PanicHandler(c)

	response, err := productCatalogFactory.NewStoreCatalogFacadeFactory().ListProducts(c.Request.Context())
	if err != nil {
		pkg.GormErrorHandler("Routes V1: ListProducts", err)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}
