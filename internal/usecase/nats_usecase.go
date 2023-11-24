package usecase

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"

	"deliveryCheck/internal/broker"
	"deliveryCheck/internal/domain"
)

type natsUsecase struct {
	orderRepository OrderRepository
}

func NewNatsUsecase(repository OrderRepository) broker.NatsUsecase {
	return &natsUsecase{
		orderRepository: repository,
	}
}

func (nu *natsUsecase) HandleMessage(m *nats.Msg) {
	log.Printf("Received a message: %s", string(m.Data))

	var order domain.Order

	err := json.Unmarshal(m.Data, &order)
	if err != nil {
		log.Printf("error: %s", err.Error())
		return
	}

	err = nu.orderRepository.Create(&order)
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}
