package repository

import (
	"time"

	"github.com/MarcoBuarque/monolito/internal/modules/checkout/domain"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Client struct {
	ID    string `gorm:"primarykey"`
	Name  string
	Email string

	// Address
	Street     string
	Number     string
	Complement string
	City       string
	State      string
	ZipCode    string
}

type Order struct {
	ID string `gorm:"primarykey"`
	// ClientID string
	Client   Client         `gorm:"foreignKey:ID"`
	Products []OrderProduct `gorm:"many2many:order_products;"`

	CreatedAt time.Time      `gorm:"->:false;column:created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"->:false;column:updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
}

type OrderProduct struct {
	OrderID   string
	ProductID string
	Price     decimal.Decimal
	Quantity  int32

	CreatedAt time.Time      `gorm:"->:false;column:created_at" json:"-"`
	UpdatedAt time.Time      `gorm:"->:false;column:updated_at" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"->:false;column:deleted_at" json:"-"`
}

func ConvertClient(data domain.ClientEntity) Client {
	return Client{
		ID:    string(data.ID().ToString()),
		Name:  data.Name(),
		Email: data.Email(),

		// Address
		Street:     data.Address().Street(),
		Number:     data.Address().Number(),
		Complement: data.Address().Complement(),
		City:       data.Address().City(),
		State:      data.Address().State(),
		ZipCode:    data.Address().ZipCode(),
	}
}

func ConvertOrderProduct(orderID string, data domain.ProductEntity) OrderProduct {
	return OrderProduct{
		OrderID:   orderID,
		ProductID: string(data.ID().ToString()),
		Price:     data.SalesPrice(),
		Quantity:  data.Quantity(),
	}
}

func Convert(data domain.OrderEntity) Order {
	products := make([]OrderProduct, len(data.Products()))
	for i := 0; i < len(data.Products()); i++ {
		products[i] = ConvertOrderProduct(data.ID().ToString(), data.Products()[i])
	}

	return Order{
		ID:       data.ID().ToString(),
		Client:   ConvertClient(data.Client()),
		Products: products,

		CreatedAt: data.CreatedAt(),
		UpdatedAt: data.UpdatedAt(),
	}
}
