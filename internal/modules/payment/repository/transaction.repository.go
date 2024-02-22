package repository

import (
	"context"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return TransactionRepository{db}
}

func (repo TransactionRepository) Save(ctx context.Context, data Transaction) error {
	if result := repo.db.WithContext(ctx).Create(&data); result.Error != nil {
		// TODO: add log
		return result.Error
	}

	return nil
}
