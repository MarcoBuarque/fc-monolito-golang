package repository

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MarcoBuarque/monolito/pkg/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	repo      ProductRepository
	mockQueue sqlmock.Sqlmock
)

func TestMain(m *testing.M) {
	db, mockQueue = database.GetDBMock()
	repo = ProductRepository{db}

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestProductRepository_NewProductRepository(t *testing.T) {
	assert := assert.New(t)

	response := NewProductRepository(db)

	assert.Equal(repo, response)
}

func TestProductRepository_FindAll(t *testing.T) {
	assert := assert.New(t)

	query := `SELECT * FROM "products" WHERE "products"."deleted_at" IS NULL`

	type args struct {
		ctx context.Context
	}

	type expect struct {
		data []ProductData
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
				ctx: context.TODO(),
			},
			setupMock: func() {
				mockQueue.ExpectQuery(query).WillReturnError(gorm.ErrRecordNotFound)
			},
			expect: expect{data: []ProductData{}, err: gorm.ErrRecordNotFound},
		},
		{
			title: "Success",
			args: args{
				ctx: context.TODO(),
			},
			setupMock: func() {
				row := sqlmock.NewRows([]string{"id", "name"}).AddRow("xpto_2", "fire").AddRow("xpto_3", "water")
				mockQueue.ExpectQuery(query).WillReturnRows(row)
			},

			expect: expect{data: []ProductData{{ID: "xpto_2", Name: "fire"}, {ID: "xpto_3", Name: "water"}}, err: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			response, err := repo.FindAll(tt.args.ctx)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data, response)
		})
	}

	assert.Nil(mockQueue.ExpectationsWereMet())
}

func TestProductRepository_Find(t *testing.T) {
	assert := assert.New(t)

	query := `SELECT * FROM "products" WHERE "products"."deleted_at" IS NULL AND "products"."id" = $1 ORDER BY "products"."id" LIMIT $2`

	type args struct {
		ctx       context.Context
		productID string
	}

	type expect struct {
		data ProductData
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
			expect:    expect{data: ProductData{}, err: fmt.Errorf("id cannot be empty")},
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
			expect: expect{data: ProductData{}, err: gorm.ErrRecordNotFound},
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

			expect: expect{data: ProductData{ID: "xpto_2", Name: "fire"}, err: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			response, err := repo.Find(tt.args.ctx, tt.args.productID)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data, response)
		})
	}

	assert.Nil(mockQueue.ExpectationsWereMet())
}
