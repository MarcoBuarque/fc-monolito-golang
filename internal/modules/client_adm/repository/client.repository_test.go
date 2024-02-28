package repository

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/MarcoBuarque/monolito/config"
	"github.com/MarcoBuarque/monolito/internal/modules/client_adm/domain"
	valueobject "github.com/MarcoBuarque/monolito/internal/modules/shared/domain/value_object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var (
	db        *gorm.DB
	repo      ClientRepository
	mockQueue sqlmock.Sqlmock
)

func TestMain(m *testing.M) {
	db, mockQueue = config.GetDBMock()
	repo = NewClientRepository(db)

	exitVal := m.Run()

	os.Exit(exitVal)
}

func TestNewClientRepository(t *testing.T) {
	assert := assert.New(t)

	response := NewClientRepository(db)

	assert.Equal(repo, response)
}

func TestClientRepository_Add(t *testing.T) {
	assert := assert.New(t)

	query := `INSERT INTO "clients" ("id","name","email","document_type","document_number","street","number","complement","city","state","zip_code","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`

	type args struct {
		ctx  context.Context
		data Client
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
				data: Client{},
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
				data: Client{Name: "xpto", Email: "email"},
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

			err := repo.Add(tt.args.ctx, tt.args.data)

			assert.Equal(tt.expect, err)
			assert.Nil(mockQueue.ExpectationsWereMet())
		})
	}
}

func TestClientRepository_Find(t *testing.T) {
	assert := assert.New(t)

	query := `SELECT * FROM "clients" WHERE "clients"."deleted_at" IS NULL AND "clients"."id" = $1 ORDER BY "clients"."id" LIMIT $2`

	type args struct {
		ctx      context.Context
		ClientID string
	}

	type expect struct {
		data Client
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
				ctx:      context.TODO(),
				ClientID: "",
			},
			setupMock: func() {},
			expect:    expect{data: Client{}, err: fmt.Errorf("id cannot be zero")},
		},
		{
			title: "Should return an error from db",
			args: args{
				ctx:      context.TODO(),
				ClientID: "57",
			},
			setupMock: func() {
				mockQueue.ExpectQuery(query).WithArgs("57", 1).WillReturnError(gorm.ErrRecordNotFound)
			},
			expect: expect{data: Client{}, err: gorm.ErrRecordNotFound},
		},
		{
			title: "Success",
			args: args{
				ctx:      context.TODO(),
				ClientID: "70",
			},
			setupMock: func() {
				row := sqlmock.NewRows([]string{"id", "name"}).AddRow(70, "fire")
				mockQueue.ExpectQuery(query).WithArgs("70", 1).WillReturnRows(row)
			},

			expect: expect{data: Client{ID: "70", Name: "fire"}, err: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			tt.setupMock()

			response, err := repo.Find(tt.args.ctx, tt.args.ClientID)
			assert.Equal(tt.expect.err, err)
			assert.Equal(tt.expect.data, response)
		})
	}

	assert.Nil(mockQueue.ExpectationsWereMet())
}

func TestConvert(t *testing.T) {
	assert := assert.New(t)

	entity, err := domain.NewClient("", "xpto", "email", valueobject.Document{}, valueobject.Address{})
	require.Nil(t, err)

	dbData := Convert(entity)

	assert.Equal(entity.ID().ToString(), dbData.ID)
	assert.Equal(entity.Name(), dbData.Name)
	assert.Equal(entity.Email(), dbData.Email)
	assert.Equal(entity.CreatedAt(), dbData.CreatedAt)
	assert.Equal(entity.UpdatedAt(), dbData.UpdatedAt)
}
