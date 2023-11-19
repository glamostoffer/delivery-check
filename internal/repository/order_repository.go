package repository

import (
	"deliveryCheck/app"
	"deliveryCheck/internal/domain"
	"fmt"
)

type orderRepository struct {
	DB app.Postgres
}

func NewOrderRepository(db app.Postgres) domain.OrderRepository {
	return &orderRepository{
		DB: db,
	}
}

func (or *orderRepository) Create(order *domain.Order) error {
	return nil
}

func (or *orderRepository) GetByUID(uid string) (domain.Order, error) {
	var order domain.Order

	query := fmt.Sprintf(`
		SELECT 
			* 
		FROM 
			"Order" 
		LEFT JOIN 
			Delivery ON "Order".delivery_id = Delivery.delivery_id 
		LEFT JOIN 
			Payment ON "Order".payment_id = Payment.payment_id 
		LEFT JOIN 
			OrderItem ON "Order".order_id = OrderItem.order_id 
		LEFT JOIN 
			Item ON OrderItem.item_id = Item.item_id 
		WHERE 
			"Order".order_uid = $1;
	`)

	err := or.DB.DB.QueryRow(query, uid).Scan(&order)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
