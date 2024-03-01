package factory

import (
	"github.com/MarcoBuarque/monolito/config"
	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/facade"
	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
	addclient "github.com/MarcoBuarque/monolito/internal/modules/client_adm/usecase/create_client"
	findclient "github.com/MarcoBuarque/monolito/internal/modules/client_adm/usecase/get_client"
)

var singleton facade.ClientAdmFacade

func NewClientAdmFacadeFactory() facade.ClientAdmFacade {
	repo := repository.NewClientRepository(config.GetDB())

	return facade.NewClientAdmFacade(
		addclient.NewCreateClientUseCase(repo),
		findclient.NewGetClientUseCase(repo),
	)
}
