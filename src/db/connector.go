package db

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
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		panic("failed to connect database")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{Conn: dbMock}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, mock
}

func setupDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	conn = db
}
