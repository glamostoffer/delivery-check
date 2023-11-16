package repository

import (
	"deliveryCheck/app"
	"deliveryCheck/internal/domain"
)

type orderRepository struct {
	db app.Postgres
}

func NewOrderRepository(db app.Postgres) domain.OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (or *orderRepository) Create(order *domain.Order) error {
	return nil
}

func (or *orderRepository) GetByUID(uid string) (domain.Order, error) {
	return domain.Order{}, nil
}
