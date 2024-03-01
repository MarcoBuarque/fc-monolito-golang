package createclient

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
)

type ICreateClientUseCase interface {
	Execute(ctx context.Context, data repository.Client) (repository.Client, error)
}
