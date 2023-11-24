package usecase

import "deliveryCheck/internal/domain"

type OrderRepository interface {
	Create(order *domain.Order) error
	GetByUID(uid string) (domain.Order, error)
}
