package getclient

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/client_adm/repository"
)

type IGetClientUseCase interface {
	Execute(ctx context.Context, clientID string) (repository.Client, error)
}
