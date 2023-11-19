package usecase

import (
	"deliveryCheck/internal/broker"
	"deliveryCheck/internal/domain"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
)

type natsUsecase struct {
	orderRepository domain.OrderRepository
}

func NewNatsUsecase(repository domain.OrderRepository) broker.NatsUsecase {
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
	}

	err = nu.orderRepository.Create(&order)
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
}
