package service

import (
	"net/http"
	"strconv"

	"github.com/FrenkenFlores/golang_l0/pkg/repository"
)

type OrderService struct {
	repository *repository.Repository
}

func (s *OrderService) GetOrder(id string) (int, map[string]any) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return http.StatusBadRequest, map[string]any{"message": "Invalid id"}
	}
	if idInt <= 0 {
		return http.StatusBadRequest, map[string]any{"message": "Id shoule be greater than 0"}
	}
	return s.repository.GetOrder(id)
}

func NewOrder(repository *repository.Repository) *OrderService {
	return &OrderService{repository: repository}
}
