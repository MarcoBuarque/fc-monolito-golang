package repository

import (
	"context"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MarcoBuarque/monolito/pkg/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	repo      TransactionRepository
	mockQueue sqlmock.Sqlmock
)

func TestMain(m *testing.M) {
	db, mockQueue = database.GetDBMock()
	repo = TransactionRepository{db}

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewTransactionRepository(t *testing.T) {
	assert := assert.New(t)

	response := NewTransactionRepository(db)

	assert.Equal(repo, response)
}

func TestTransactionRepository_Add(t *testing.T) {
	assert := assert.New(t)

	query := `INSERT INTO "transactions" ("id","order_id","amount","status","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7)`

	type args struct {
		ctx  context.Context
		data Transaction
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
				data: Transaction{},
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
				data: Transaction{},
			},
			setupMock: func() {
				mockQueue.ExpectBegin()
				mockQueue.ExpectExec(query).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
				mockQueue.ExpectCommit()
			},
			expect: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			err := repo.Save(tt.args.ctx, tt.args.data)

			assert.Equal(tt.expect, err)
		})
	}

	assert.Nil(mockQueue.ExpectationsWereMet())
}
