package v1

import (
	"net/http"

	"github.com/MarcoBuarque/monolito/constant"
	clientAdmFactory "github.com/MarcoBuarque/monolito/internal/modules/client_adm/factory"
	repoClientADM "github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
	"github.com/MarcoBuarque/monolito/pkg"
	"github.com/gin-gonic/gin"
)

// CreateClient godoc
//
//	@Summary		Create a new cleint.
//	@Description	This endpoint allows you to create a new client by providing its details in the request body.
//	@Tags			clients
//	@Accept			json
//	@Produce		json
//	@Param			client	body		repoClientADM.Client	true	"Client data"
//	@Success		200		{object}	pkg.ApiResponse[repoClientADM.Client]
//
//	@Failure		400		{object}	pkg.ApiResponse[pkg.Null]
//	@Failure		500		{object}	pkg.ApiResponse[pkg.Null]
//
//	@Router			/client-management [post]
func CreateClient(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request repoClientADM.Client
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	response, err := clientAdmFactory.NewClientAdmFacadeFactory().Add(c.Request.Context(), request)
	if err != nil {
		pkg.GormErrorHandler("Routes V1: CreateClient", err)
	}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, response))
}

// GetClient godoc
//
//	@Summary		Retrieve a specific cleint by its ID.
//	@Description	This endpoint retrieves the details of a single cleint based on the provided cleint ID.
//	@Tags			clients
//	@Accept			json
//	@Produce		json
//
//	@Param			clientID	path		string	true	"cleint ID"
//
//	@Success		200			{object}	pkg.ApiResponse[repoClientADM.Client]
//
//	@Failure		400			{object}	pkg.ApiResponse[pkg.Null]
//	@Failure		500			{object}	pkg.ApiResponse[pkg.Null]
//
//	@Router			/client-management/{clientID} [get]
func GetClient(c *gin.Context) {
	defer pkg.PanicHandler(c)

	clientID := c.Param("clientID")

	response, err := clientAdmFactory.NewClientAdmFacadeFactory().Find(c.Request.Context(), clientID)
	if err != nil {
		pkg.GormErrorHandler("Routes V1: GetClient", err)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}
