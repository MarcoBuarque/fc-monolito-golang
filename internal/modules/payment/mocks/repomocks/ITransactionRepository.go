// Code generated by mockery v1.0.0. DO NOT EDIT.

package repomocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import repository "github.com/MarcoBuarque/monolito/internal/modules/payment/repository"

// ITransactionRepository is an autogenerated mock type for the ITransactionRepository type
type ITransactionRepository struct {
	mock.Mock
}

// Save provides a mock function with given fields: ctx, data
func (_m *ITransactionRepository) Save(ctx context.Context, data repository.Transaction) error {
	ret := _m.Called(ctx, data)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, repository.Transaction) error); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}