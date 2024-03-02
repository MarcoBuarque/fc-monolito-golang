package clientmanagement

import (
	"net/http"

	"github.com/MarcoBuarque/fc-monolito-golang/config"
	"github.com/MarcoBuarque/fc-monolito-golang/constant"
	repository "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/repository"
	createclient "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/usecase/create_client"
	getclient "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/usecase/get_client"
	"github.com/MarcoBuarque/fc-monolito-golang/pkg"
	"github.com/gin-gonic/gin"
)

var (
	createUseCase createclient.ICreateClientUseCase
	getUseCase    getclient.IGetClientUseCase
)

func init() {
	repo := repository.NewClientRepository(config.GetDB())

	createUseCase = createclient.NewCreateClientUseCase(repo)
	getUseCase = getclient.NewGetClientUseCase(repo)
}

func ConfigRoutes(v1 *gin.RouterGroup) {
	clientManagement := v1.Group("/client-management")
	{
		clientManagement.POST("/", createClient)
		// clientManagement.GET("", )
		clientManagement.GET("/:clientID", getClient)
	}
}

// createClient godoc
//
//	@Summary		Create a new cleint.
//	@Description	This endpoint allows you to create a new client by providing its details in the request body.
//	@Tags			clients
//	@Accept			json
//	@Produce		json
//	@Param			client	body		repository.Client	true	"Client data"
//	@Success		200		{object}	pkg.ApiResponse[repository.Client]
//
//	@Failure		400		{object}	pkg.ApiResponse[pkg.Null]
//	@Failure		500		{object}	pkg.ApiResponse[pkg.Null]
//
//	@Router			/client-management [post]
func createClient(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request repository.Client
	if err := c.ShouldBindJSON(&request); err != nil {
		pkg.PanicException(constant.InvalidRequest)
	}

	response, err := createUseCase.Execute(c.Request.Context(), request)
	if err != nil {
		pkg.GormErrorHandler("Routes V1: CreateClient:", err)
	}

	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, response))
}

// getClient godoc
//
//	@Summary		Retrieve a specific cleint by its ID.
//	@Description	This endpoint retrieves the details of a single cleint based on the provided cleint ID.
//	@Tags			clients
//	@Accept			json
//	@Produce		json
//
//	@Param			clientID	path		string	true	"cleint ID"
//
//	@Success		200			{object}	pkg.ApiResponse[repository.Client]
//
//	@Failure		400			{object}	pkg.ApiResponse[pkg.Null]
//	@Failure		500			{object}	pkg.ApiResponse[pkg.Null]
//
//	@Router			/client-management/{clientID} [get]
func getClient(c *gin.Context) {
	defer pkg.PanicHandler(c)

	clientID := c.Param("clientID")

	response, err := getUseCase.Execute(c.Request.Context(), clientID)
	if err != nil {
		pkg.GormErrorHandler("Routes V1: GetClient:", err)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, response))
}
