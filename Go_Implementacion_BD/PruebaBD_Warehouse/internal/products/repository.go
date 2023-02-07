package products

import (
	"PruebaBD_Warehouse/internal/domain"
	"errors"
)

var(
	ErrNotFound = errors.New("product not found")
	ErrInternal = errors.New("an internal error")
	ErrDuplicated = errors.New("duplicated product")
)

type Repository interface{
	Get(id int) (*domain.Product, error)
	Store(product *domain.Product) error
	Update(product *domain.Product) error
	Delete(id int) error
}
