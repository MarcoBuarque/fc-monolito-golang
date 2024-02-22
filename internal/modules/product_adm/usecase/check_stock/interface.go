package checkstock

import "context"

type ICheckStockUseCase interface {
	Execute(ctx context.Context, id string) (int32, error)
}
