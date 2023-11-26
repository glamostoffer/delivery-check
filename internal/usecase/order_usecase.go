package usecase

import (
	"context"
	"deliveryCheck/internal/domain"
	"deliveryCheck/internal/repository"
)

type orderUsecase struct {
	orderRepository repository.OrderRepository
}

func NewOrderUsecase(repository repository.OrderRepository) OrderUsecase {
	return &orderUsecase{
		orderRepository: repository,
	}
}

func (ou *orderUsecase) GetOrderByUID(ctx context.Context, uid string) (domain.Order, error) {
	return ou.orderRepository.GetByUID(uid)
}
