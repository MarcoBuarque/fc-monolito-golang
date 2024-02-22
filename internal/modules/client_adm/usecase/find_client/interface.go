package findclient

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
)

type IFindClientUseCase interface {
	Execute(ctx context.Context, clientID string) (repository.Client, error)
}
