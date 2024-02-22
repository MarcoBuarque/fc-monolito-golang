package repository

import "context"

type ITransactionRepository interface {
	Save(ctx context.Context, data Transaction) error
}
