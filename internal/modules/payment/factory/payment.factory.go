package factory

import (
	"github.com/MarcoBuarque/fc-monolito-golang/config"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/payment/facade"
	"github.com/MarcoBuarque/fc-monolito-golang/internal/modules/payment/repository"
	processpayment "github.com/MarcoBuarque/fc-monolito-golang/internal/modules/payment/usecase/process_payment"
)

type PaymentFacadeFactory struct {
}

func NewPaymentFacadeFactory() facade.PaymentFacade {
	repo := repository.NewTransactionRepository(config.GetDB())

	return facade.NewPaymentFacade(
		processpayment.NewProcessPaymentUseCase(repo),
	)
}
