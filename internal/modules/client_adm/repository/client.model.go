package repository

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	ID        string `gorm:"primarykey"`
	Name      string
	Email     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
