package repository

import gol0 "github.com/FrenkenFlores/golang_l0"

type Repository struct {
	Order
}

type Order interface {
	GetOrder(id string) (int, map[string]any)
	SetOrder(orderObj gol0.Order)
}

func NewRepository(cfg Config) (*Repository, error) {
	order, err := NewPostgresDb(cfg)
	return &Repository{
		Order: order,
	}, err
}
