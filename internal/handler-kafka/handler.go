package handlerkafka

import (
	"encoding/json"

	gol0 "github.com/FrenkenFlores/golang_l0"
	"github.com/FrenkenFlores/golang_l0/pkg/repository"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/sirupsen/logrus"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleMessage(message []byte, offset kafka.Offset, repo repository.Repository) error {
	order := gol0.Order{}
	err := json.Unmarshal(message, &order)
	if err != nil {
		logrus.Error(err)
		return err
	}
	repo.SetOrder(order)
	return nil
}
