package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	conn *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	once.Do(setupDB)

	return conn
}

func GetDBMock() (*gorm.DB, sqlmock.Sqlmock) {
	dbMock, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic("failed to start sqlmock")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: dbMock}), &gorm.Config{})
	if err != nil {
		panic("failed to connect mocked database")
	}

	return db, mock
}

func setupDB() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:123456@db:5432/fc-monolito?sslmode=prefer"
	}

	fmt.Println("LOCALLL", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{CreateBatchSize: 100})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err))
	}

	conn = db
}
