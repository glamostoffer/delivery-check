package usecase

import (
	"context"
	"deliveryCheck/internal/domain"
)

type OrderUsecase interface {
	GetOrderByUID(ctx context.Context, uid string) (domain.Order, error)
}
