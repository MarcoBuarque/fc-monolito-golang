package factory

import (
	"sync"

	"github.com/MarcoBuarque/fc-monolito-golang/config"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/facade"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/repository"
	createclient "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/usecase/create_client"
	getclient "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/usecase/get_client"
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
