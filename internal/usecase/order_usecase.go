package usecase

import (
	"context"
	"deliveryCheck/internal/domain"
)

type orderUsecase struct {
	orderRepository domain.OrderRepository
}

func NewOrderUsecase(repository domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{
		orderRepository: repository,
	}
}

func (ou *orderUsecase) GetOrderByUID(ctx context.Context, uid string) (domain.Order, error) {
	return domain.Order{}, nil
}
