package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
)

type IClientAdmFacade interface {
	CreateClient(ctx context.Context, data repository.Client) (repository.Client, error)
	GetClient(ctx context.Context, clientID string) (repository.Client, error)
}
