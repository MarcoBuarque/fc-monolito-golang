package processpayment

import (
	"context"

	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/payment/domain"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/payment/repository"
	"github.com/shopspring/decimal"
)

type ProcessPaymentUseCase struct {
	transactionRepository repository.ITransactionRepository
}

func NewProcessPaymentUseCase(repo repository.ITransactionRepository) ProcessPaymentUseCase {
	return ProcessPaymentUseCase{repo}
}

func (controller ProcessPaymentUseCase) Execute(ctx context.Context, orderID string, amount decimal.Decimal) (repository.Transaction, error) {
	txEntity, err := domain.NewTransaction("", orderID, "", amount)
	if err != nil {
		// TODO: add log
		return repository.Transaction{}, err
	}

	data := repository.Convert(txEntity)

	if err := controller.transactionRepository.Save(ctx, data); err != nil {
		// TODO: add log
		return repository.Transaction{}, err
	}

	return data, nil
}
