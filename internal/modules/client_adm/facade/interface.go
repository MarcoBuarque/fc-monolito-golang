package facade

import (
	"context"

	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/repository"
)

type IClientAdmFacade interface {
	Add(ctx context.Context, data repository.Client) (repository.Client, error)
	Find(ctx context.Context, ClientID string) (repository.Client, error)
}
