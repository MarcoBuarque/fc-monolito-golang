package v1

import (
	"net/http"

	productAdmFactory "github.com/MarcoBuarque/monolito/internal/modules/product_adm/factory"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"
	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var data repository.Product
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	response, err := productAdmFactory.NewProductAdmFacadeFactory().AddProduct(c.Request.Context(), data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Product saved successfully", "data": response})
}
