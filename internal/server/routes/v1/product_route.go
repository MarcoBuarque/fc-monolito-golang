package v1

import (
	"net/http"

	"github.com/MarcoBuarque/monolito/constant"
	productAdmFactory "github.com/MarcoBuarque/monolito/internal/modules/product_adm/factory"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	"github.com/MarcoBuarque/monolito/internal/modules/store_catalog/factory"
	"github.com/MarcoBuarque/monolito/pkg"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var data repository.Product
	if err := c.ShouldBindJSON(&data); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	response, err := productAdmFactory.NewProductAdmFacadeFactory().CreateProduct(c.Request.Context(), data)
	if err != nil {
		pkg.GormErrorHandler("Routes V1: CreateProduct", err)
	}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, response))
}

func GetProduct(c *gin.Context) {
	defer pkg.PanicHandler(c)

	productID := c.Param("productID")

	response, err := factory.NewStoreCatalogFacadeFactory().GetProduct(c.Request.Context(), productID)
	if err != nil {
		pkg.GormErrorHandler("Routes V1: GetProduct", err)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

func ListProducts(c *gin.Context) {
	defer pkg.PanicHandler(c)

	response, err := factory.NewStoreCatalogFacadeFactory().ListProducts(c.Request.Context())
	if err != nil {
		pkg.GormErrorHandler("Routes V1: ListProducts", err)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}

func UpdateSalesPrice(c *gin.Context) {
	defer pkg.PanicHandler(c)

	productID := c.Param("productID")

	var data repository.Product
	if err := c.ShouldBindJSON(&data); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	parserPrice, ok := data.SalesPrice.Float64()
	if !ok {
		pkg.PanicException(constant.InvalidRequest)
	}

	err := factory.NewStoreCatalogFacadeFactory().UpdateSalesPrice(c.Request.Context(), productID, parserPrice)
	if err != nil {
		pkg.GormErrorHandler("Routes V1: ListProducts", err)
	}

	// TODO: Review: send all product data back?
	c.JSON(http.StatusNoContent, pkg.BuildResponse(constant.Success, pkg.Null()))
}
