package config

import (
	"sync"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
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
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{CreateBatchSize: 100})
	if err != nil {
		panic("failed to connect database")
	}

	conn = db
}
