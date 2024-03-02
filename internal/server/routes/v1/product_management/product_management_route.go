package productmanagement

import (
	"net/http"

	"github.com/MarcoBuarque/fc-monolito-golang/config"
	"github.com/MarcoBuarque/fc-monolito-golang/constant"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/repository"
	createproduct "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/usecase/create_product"
	updatesalesprice "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/product_adm/usecase/update_sales_price"
	"github.com/MarcoBuarque/fc-monolito-golang/pkg"
	"github.com/gin-gonic/gin"
)

type updateSalesPriceRequest struct {
	Price float32 `binding:"required"`
}

var (
	createUsecase      createproduct.ICreateProductUseCase
	updatePriceUsecase updatesalesprice.IUpdateSalesPriceUseCase
)

func init() {
	repo := repository.NewProductRepository(config.GetDB())

	createUsecase = createproduct.NewCreateProductUseCase(repo)
	updatePriceUsecase = updatesalesprice.NewUpdateSalesPriceUseCase(repo)
}

func ConfigRoutes(v1 *gin.RouterGroup) {
	productManagement := v1.Group("/product-management")
	{
		productManagement.POST("/", createProduct)
		productManagement.PATCH("/:productID", updateSalesPrice)
	}
}

// createProduct godoc
//
//	@Summary		Create a new product.
//	@Description	This endpoint allows you to create a new product by providing its details in the request body.
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			product	body		repository.Product	true	"Product data"
//	@Success		200		{object}	pkg.ApiResponse[repository.Product]
//
//	@Failure		400		{object}	pkg.ApiResponse[pkg.Null]
//	@Failure		500		{object}	pkg.ApiResponse[pkg.Null]
//
//	@Router			/product-management [post]
func createProduct(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request repository.Product
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	response, err := createUsecase.Execute(c.Request.Context(), request)
	if err != nil {
		pkg.GormErrorHandler("Routes V1: CreateProduct:", err)
	}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, response))
}

// updateSalesPrice godoc
//
//	@Summary		Update the sales price of a product.
//	@Description	This API allows you to update the sales price of a specific product by providing its ID and the new price in the request body.
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			productID	path	string					true	"Product ID"
//	@Param			product		body	updateSalesPriceRequest	true	"new price value"
//
// @Success	200	{object}	pkg.ApiResponse[repository.Product]
//
// @Failure	400	{object}	pkg.ApiResponse[pkg.Null]
// @Failure	500	{object}	pkg.ApiResponse[pkg.Null]
//
// @Router		/product-management/{productID} [patch]
func updateSalesPrice(c *gin.Context) {
	defer pkg.PanicHandler(c)

	productID := c.Param("productID")

	var request updateSalesPriceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	response, err := updatePriceUsecase.Execute(c.Request.Context(), productID, request.Price)
	if err != nil {
		pkg.GormErrorHandler("Routes V1: UpdateSalesPrice:", err)
	}

	// TODO: Review: send all product data back?
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}
