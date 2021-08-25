package customers

import (
	"context"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName   string `gorm:"size:50;not null;" json:"first_name"`
	LastName    string `gorm:"size:50;not null;" json:"last_name"`
	PhoneNumber string `gorm:"size:16;not null;" json:"phone_number"`
	Address     string `gorm:"size:99;not null;" json:"address"`
	City        string `gorm:"size:50;not null;" json:"city"`
	State       string `gorm:"size:25;not null;" json:"state"`
	ZipCode     string `gorm:"size:10;not null;" json:"zip_code"`
}

type Repository interface {
	GetCustomerById(ctx context.Context, id uint) (Customer, error)
	GetAllCustomers(ctx context.Context) ([]Customer, error)
	CreateCustomer(ctx context.Context, customer Customer) (string, error)
	UpdateCustomer(ctx context.Context, customer Customer) (string, error)
	DeleteCustomer(ctx context.Context, id uint) error
}
