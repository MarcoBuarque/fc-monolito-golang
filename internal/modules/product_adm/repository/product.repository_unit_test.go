package repository

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MarcoBuarque/monolito/config"
	"github.com/MarcoBuarque/monolito/internal/modules/product_adm/domain"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	repo      ProductRepository
	mockQueue sqlmock.Sqlmock
)

func TestMain(m *testing.M) {
	db, mockQueue = config.GetDBMock()
	repo = NewProductRepository(db)

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestProductRepository_NewProductRepository(t *testing.T) {
	assert := assert.New(t)

	response := NewProductRepository(db)

	assert.Equal(repo, response)
}

func TestProductRepository_CreateProduct(t *testing.T) {
	assert := assert.New(t)

	query := `INSERT INTO "products" ("id","name","description","purchase_price","sales_price","stock","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	type args struct {
		ctx  context.Context
		data Product
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
				data: Product{},
			},
			setupMock: func() {
				mockQueue.ExpectBegin()
				mockQueue.ExpectExec(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnError(gorm.ErrRecordNotFound)
				mockQueue.ExpectRollback()
			},
			expect: gorm.ErrRecordNotFound,
		},
		{
			title: "Success",
			args: args{
				ctx:  context.TODO(),
				data: Product{},
			},
			setupMock: func() {
				mockQueue.ExpectBegin()
				mockQueue.ExpectExec(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
				mockQueue.ExpectCommit()
			},
			expect: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			err := repo.CreateProduct(tt.args.ctx, tt.args.data)

			assert.Equal(tt.expect, err)
		})
	}

	assert.Nil(mockQueue.ExpectationsWereMet())
}

func TestProductRepository_GetProduct(t *testing.T) {
	assert := assert.New(t)

	query := `SELECT * FROM "products" WHERE "products"."deleted_at" IS NULL AND "products"."id" = $1 ORDER BY "products"."id" LIMIT $2`

	type args struct {
		ctx       context.Context
		productID string
	}

	type expect struct {
		data Product
		err  error
	}

	tests := []struct {
		title     string
		args      args
		setupMock func()
		expect    expect
	}{
		{
			title: "Should return an error from db",
			args: args{
				ctx:       context.TODO(),
				productID: "",
			},
			setupMock: func() {},
			expect:    expect{data: Product{}, err: fmt.Errorf("id cannot be empty")},
		},
		{
			title: "Should return an error from db",
			args: args{
				ctx:       context.TODO(),
				productID: "xpto",
			},
			setupMock: func() {
				mockQueue.ExpectQuery(query).WithArgs("xpto", 1).WillReturnError(gorm.ErrRecordNotFound)
			},
			expect: expect{data: Product{}, err: gorm.ErrRecordNotFound},
		},
		{
			title: "Success",
			args: args{
				ctx:       context.TODO(),
				productID: "xpto_2",
			},
			setupMock: func() {
				row := sqlmock.NewRows([]string{"id", "name"}).AddRow("xpto_2", "fire")
				mockQueue.ExpectQuery(query).WithArgs("xpto_2", 1).WillReturnRows(row)
			},

			expect: expect{data: Product{ID: "xpto_2", Name: "fire"}, err: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			response, err := repo.GetProduct(tt.args.ctx, tt.args.productID)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data, response)
		})
	}

	assert.Nil(mockQueue.ExpectationsWereMet())
}

func TestConvert(t *testing.T) {
	assert := assert.New(t)

	entity, err := domain.NewProduct("", "xpto", "xpto_description", decimal.NewFromInt(10), 10)
	require.Nil(t, err)

	dbData := Convert(entity)

	assert.Equal(entity.ID().ToString(), dbData.ID)
	assert.Equal(entity.Name(), dbData.Name)
	assert.Equal(entity.PurchasePrice(), dbData.PurchasePrice)
	assert.Equal(entity.PurchasePrice(), dbData.SalesPrice)
	assert.Equal(entity.Stock(), dbData.Stock)
	assert.Equal(entity.CreatedAt(), dbData.CreatedAt)
	assert.Equal(entity.UpdatedAt(), dbData.UpdatedAt)
}
