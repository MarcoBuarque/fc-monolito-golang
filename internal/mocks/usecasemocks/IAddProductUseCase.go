// Code generated by mockery v1.0.0. DO NOT EDIT.

package usecasemocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import repository "github.com/MarcoBuarque/monolito/internal/modules/product_adm/repository"

// IAddProductUseCase is an autogenerated mock type for the IAddProductUseCase type
type IAddProductUseCase struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, data
func (_m *IAddProductUseCase) Execute(ctx context.Context, data repository.ProductData) (repository.ProductData, error) {
	ret := _m.Called(ctx, data)

	var r0 repository.ProductData
	if rf, ok := ret.Get(0).(func(context.Context, repository.ProductData) repository.ProductData); ok {
		r0 = rf(ctx, data)
	} else {
		r0 = ret.Get(0).(repository.ProductData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repository.ProductData) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
