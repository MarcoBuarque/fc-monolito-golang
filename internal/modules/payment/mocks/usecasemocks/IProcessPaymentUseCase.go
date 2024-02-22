// Code generated by mockery v1.0.0. DO NOT EDIT.

package usecasemocks

import context "context"
import decimal "github.com/shopspring/decimal"
import mock "github.com/stretchr/testify/mock"

import repository "github.com/MarcoBuarque/monolito/internal/modules/payment/repository"

// IProcessPaymentUseCase is an autogenerated mock type for the IProcessPaymentUseCase type
type IProcessPaymentUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, orderID, amount
func (_m *IProcessPaymentUseCase) Execute(ctx context.Context, orderID string, amount decimal.Decimal) (repository.Transaction, error) {
	ret := _m.Called(ctx, orderID, amount)

	var r0 repository.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, string, decimal.Decimal) repository.Transaction); ok {
		r0 = rf(ctx, orderID, amount)
	} else {
		r0 = ret.Get(0).(repository.Transaction)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, decimal.Decimal) error); ok {
		r1 = rf(ctx, orderID, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}