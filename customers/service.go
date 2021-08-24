package customers

import "context"

type Service interface {
	GetById(ctx context.Context, id string) (Customer, error)
	GetAll(ctx context.Context) ([]Customer, error)
	Create(ctx context.Context, customer Customer) (string, error)
	Update(ctx context.Context, customer Customer) (string, error)
	Delete(ctx context.Context, id string) error
}
