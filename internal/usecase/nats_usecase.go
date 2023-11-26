package usecase

import (
	"deliveryCheck/internal/repository"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"

	"deliveryCheck/internal/broker"
	"deliveryCheck/internal/domain"
)

type natsUsecase struct {
	orderRepository repository.OrderRepository
}

func NewNatsUsecase(repository repository.OrderRepository) broker.NatsUsecase {
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
