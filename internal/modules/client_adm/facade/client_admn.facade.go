package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
	addclient "github.com/MarcoBuarque/monolito/internal/modules/client_adm/usecase/add_client"
	findclient "github.com/MarcoBuarque/monolito/internal/modules/client_adm/usecase/find_client"
)

type ClientAdmFacade struct {
	addUseCase  addclient.IAddClientUseCase
	findUseCase findclient.IFindClientUseCase
}

func NewClientAdmFacade(findAllUseCase addclient.IAddClientUseCase, findUseCase findclient.IFindClientUseCase) ClientAdmFacade {
	return ClientAdmFacade{findAllUseCase, findUseCase}
}

func (facade ClientAdmFacade) Add(ctx context.Context, data repository.Client) (repository.Client, error) {
	return facade.addUseCase.Execute(ctx, data)
}

func (facade ClientAdmFacade) Find(ctx context.Context, clientID string) (repository.Client, error) {
	return facade.findUseCase.Execute(ctx, clientID)
}
