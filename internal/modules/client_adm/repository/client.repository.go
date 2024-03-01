package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type ClientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return ClientRepository{db}
}

func (repo ClientRepository) CreateClient(ctx context.Context, data Client) error {
	if result := repo.db.WithContext(ctx).Create(&data); result.Error != nil {
		// TODO: add log
		return result.Error
	}

	return nil
}

func (repo ClientRepository) GetClient(ctx context.Context, clientID string) (Client, error) {
	if clientID == "" {
		return Client{}, fmt.Errorf("id cannot be zero")
	}

	response := &Client{ID: clientID}
	if result := repo.db.WithContext(ctx).First(response); result.Error != nil {
		// TODO: add log
		return Client{}, result.Error
	}

	return *response, nil
}
