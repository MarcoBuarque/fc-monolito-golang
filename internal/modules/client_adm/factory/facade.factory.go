package factory

import (
	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/facade"
	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
	addclient "github.com/MarcoBuarque/monolito/internal/modules/client_adm/usecase/add_client"
	findclient "github.com/MarcoBuarque/monolito/internal/modules/client_adm/usecase/find_client"
	"github.com/MarcoBuarque/monolito/pkg/database"
)

func NewClientAdmFacadeFactory() facade.ClientAdmFacade {
	repo := repository.NewClientRepository(database.GetDB())

	return facade.NewClientAdmFacade(
		addclient.NewAddClientUseCase(repo),
		findclient.NewFindClientUseCase(repo),
	)
}
