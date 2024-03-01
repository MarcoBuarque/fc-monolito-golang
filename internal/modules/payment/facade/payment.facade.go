package facade

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/payment/repository"
	processpayment "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/payment/usecase/process_payment"
	"github.com/shopspring/decimal"
)

type PaymentFacade struct {
	processPaymentUseCase processpayment.IProcessPaymentUseCase
}

func NewPaymentFacade(processPayUseCase processpayment.IProcessPaymentUseCase) PaymentFacade {
	return PaymentFacade{processPayUseCase}
}

func (facade PaymentFacade) Process(ctx context.Context, orderID string, amount decimal.Decimal) (repository.Transaction, error) {
	return facade.processPaymentUseCase.Execute(ctx, orderID, amount)
}
