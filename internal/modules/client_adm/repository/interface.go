package repository

import "context"

type IClientRepository interface {
	CreateClient(ctx context.Context, data Client) error
	GetClient(ctx context.Context, clientID string) (Client, error)
}
