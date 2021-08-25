package customers

import (
	"context"
	"database/sql"
)

var _ Repository = (*Repos)(nil)

type Repos struct {
	Db *sql.DB
}

func (r *Repos) GetCustomerById(ctx context.Context, id uint) (Customer, error) {
	panic("implement me")
}

func (r *Repos) GetAllCustomers(ctx context.Context) ([]Customer, error) {
	panic("implement me")
}

func (r *Repos) CreateCustomer(ctx context.Context, customer Customer) (string, error) {
	panic("implement me")
}

func (r *Repos) UpdateCustomer(ctx context.Context, customer Customer) (string, error) {
	panic("implement me")
}

func (r *Repos) DeleteCustomer(ctx context.Context, id uint) error {
	panic("implement me")
}



