package facade

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/repository"
	createclient "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/usecase/create_client"
	getclient "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/usecase/get_client"
)

type ClientAdmFacade struct {
	createUseCase createclient.ICreateClientUseCase
	getUseCase    getclient.IGetClientUseCase
}

func NewClientAdmFacade(listUsersUseCase createclient.ICreateClientUseCase, getUseCase getclient.IGetClientUseCase) ClientAdmFacade {
	return ClientAdmFacade{listUsersUseCase, getUseCase}
}

func (facade ClientAdmFacade) CreateClient(ctx context.Context, data repository.Client) (repository.Client, error) {
	return facade.createUseCase.Execute(ctx, data)
}

func (facade ClientAdmFacade) GetClient(ctx context.Context, clientID string) (repository.Client, error) {
	return facade.getUseCase.Execute(ctx, clientID)
}
