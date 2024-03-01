package factory

import (
	"sync"

	"github.com/MarcoBuarque/monolito/config"
	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/facade"
	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
	createclient "github.com/MarcoBuarque/monolito/internal/modules/client_adm/usecase/create_client"
	getclient "github.com/MarcoBuarque/monolito/internal/modules/client_adm/usecase/get_client"
)

var (
	once      sync.Once
	singleton facade.ClientAdmFacade
)

func createSingleton() {
	repo := repository.NewClientRepository(config.GetDB())

	singleton = facade.NewClientAdmFacade(
		createclient.NewCreateClientUseCase(repo),
		getclient.NewGetClientUseCase(repo),
	)
}

func NewClientAdmFacadeFactory() facade.ClientAdmFacade {

	once.Do(createSingleton)

	return singleton
}
