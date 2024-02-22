package addclient

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
)

type IAddClientUseCase interface {
	Execute(ctx context.Context, data repository.Client) (repository.Client, error)
}
