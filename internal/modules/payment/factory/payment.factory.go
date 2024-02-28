package factory

import (
	"github.com/MarcoBuarque/monolito/config"
	"github.com/MarcoBuarque/monolito/internal/modules/payment/facade"
	"github.com/MarcoBuarque/monolito/internal/modules/payment/repository"
	processpayment "github.com/MarcoBuarque/monolito/internal/modules/payment/usecase/process_payment"
)

type PaymentFacadeFactory struct {
}

func NewPaymentFacadeFactory() facade.PaymentFacade {
	repo := repository.NewTransactionRepository(config.GetDB())

	return facade.NewPaymentFacade(
		processpayment.NewProcessPaymentUseCase(repo),
	)
}
