package usecase

import (
	"context"
	"deliveryCheck/internal/domain"
)

type orderUsecase struct {
	orderRepository OrderRepository
}

func NewOrderUsecase(repository OrderRepository) domain.OrderUsecase {
	return &orderUsecase{
		orderRepository: repository,
	}
}

func (ou *orderUsecase) GetOrderByUID(ctx context.Context, uid string) (domain.Order, error) {
	return ou.orderRepository.GetByUID(uid)
}
