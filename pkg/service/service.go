package service

import "github.com/FrenkenFlores/golang_l0/pkg/repository"

type Service struct {
	Order
}

type Order interface {
	GetOrder(id string) (int, map[string]any)
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Order: NewOrder(repository),
	}
}
