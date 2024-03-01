package getclient

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
)

type IGetClientUseCase interface {
	Execute(ctx context.Context, clientID string) (repository.Client, error)
}
