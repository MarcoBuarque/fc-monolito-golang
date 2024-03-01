package facade

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/payment/repository"
	"github.com/shopspring/decimal"
)

type IPaymentFacade interface {
	Process(ctx context.Context, orderID string, amount decimal.Decimal) (repository.Transaction, error)
}
