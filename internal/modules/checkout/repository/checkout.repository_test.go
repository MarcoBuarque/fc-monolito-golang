package repository

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MarcoBuarque/monolito/config"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	repo      CheckoutRepository
	mockQueue sqlmock.Sqlmock
)

func TestMain(m *testing.M) {
	db, mockQueue = config.GetDBMock()
	repo = NewCheckoutRepository(db)

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewCheckoutRepository(t *testing.T) {
	assert := assert.New(t)

	response := NewCheckoutRepository(db)

	assert.Equal(repo, response)
}

func TestCheckoutRepository_AddOrder(t *testing.T) {
	assert := assert.New(t)

	order := Order{
		Client: Client{ID: "XPTO"},
		Products: []OrderProduct{
			{
				OrderID:   "order_id",
				ProductID: "product_id",
				Price:     decimal.NewFromFloat32(10),
				Quantity:  5,
			},
			{
				OrderID:   "order_id_2",
				ProductID: "product_id_2",
				Price:     decimal.NewFromFloat32(7),
				Quantity:  2,
			},
		},
	}

	query := ``

	type args struct {
		ctx  context.Context
		data Order
	}

	tests := []struct {
		title     string
		args      args
		setupMock func()
		expect    error
	}{
		{
			title: "Should return an error from db",
			args: args{
				ctx:  context.TODO(),
				data: Order{},
			},
			setupMock: func() {
				mockQueue.ExpectBegin().WillReturnError(gorm.ErrInvalidDB)
			},
			expect: gorm.ErrInvalidDB,
		},
		{
			title: "Success",
			args: args{
				ctx:  context.TODO(),
				data: order,
			},
			setupMock: func() {
				mockQueue.ExpectBegin()
				mockQueue.ExpectExec(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
				mockQueue.ExpectCommit()
			},
			expect: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			err := repo.AddOrder(tt.args.ctx, tt.args.data)

			assert.Equal(tt.expect, err)
			assert.Nil(mockQueue.ExpectationsWereMet())
		})
	}
}

func TestCheckoutRepository_FindOrder(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx     context.Context
		orderID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := CheckoutRepository{
				db: tt.fields.db,
			}
			got, err := repo.FindOrder(tt.args.ctx, tt.args.orderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckoutRepository.FindOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckoutRepository.FindOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
