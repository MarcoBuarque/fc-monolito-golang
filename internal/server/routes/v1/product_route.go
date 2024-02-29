package v1

import (
	"net/http"

	"github.com/MarcoBuarque/monolito/constant"
	productAdmFactory "github.com/MarcoBuarque/monolito/internal/modules/product_adm/factory"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
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
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, response))
}

func GetProduct(c *gin.Context) {

}
