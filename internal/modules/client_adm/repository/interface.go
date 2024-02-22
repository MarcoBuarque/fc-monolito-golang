package repository

import "context"

type IClientRepository interface {
	Add(ctx context.Context, data Client) error
	Find(ctx context.Context, clientID string) (Client, error)
}
