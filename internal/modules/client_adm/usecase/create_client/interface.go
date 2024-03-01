package createclient

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/repository"
)

type ICreateClientUseCase interface {
	Execute(ctx context.Context, data repository.Client) (repository.Client, error)
}
