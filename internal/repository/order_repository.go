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
	var deliveryId int
	query := fmt.Sprintf("INSERT INTO delivery (name, phone, zip, city, address, region, email) values ($1, $2, $3, $4, $5, $6, $7) RETURNING delivery_id")
	err := or.DB.DB.QueryRow(query, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email).Scan(&deliveryId)
	if err != nil {
		return err
	}

	var paymentId int
	query = fmt.Sprintf("INSERT INTO payment (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING payment_id")
	err = or.DB.DB.QueryRow(query, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDT, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee).Scan(&paymentId)
	if err != nil {
		return err
	}

	var orderId int
	query = fmt.Sprintf("INSERT INTO \"Order\" (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING order_id")
	err = or.DB.DB.QueryRow(query, order.OrderUID, order.TrackNumber, order.Entry, deliveryId, paymentId, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SMID, order.DateCreated, order.OOFShard).Scan(&orderId)
	if err != nil {
		return err
	}

	for _, item := range order.Items {
		var itemId int
		query = fmt.Sprintf("INSERT INTO item (order_id, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING item_id")
		err = or.DB.DB.QueryRow(query, orderId, item.ChrtID, item.TrackNumber, item.Price, item.RID, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status).Scan(&itemId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (or *orderRepository) GetByUID(uid string) (domain.Order, error) {
	var orderId int
	var deliveryId int
	var paymentId int
	query := fmt.Sprintf(`SELECT order_id, delivery_id, payment_id FROM  "Order" WHERE order_uid = $1`)
	err := or.DB.DB.QueryRow(query, uid).Scan(&orderId, &deliveryId, &paymentId)
	if err != nil {
		return domain.Order{}, nil
	}

	var delivery domain.Delivery
	query = fmt.Sprintf("SELECT name, phone, zip, city, address, region, email FROM  delivery WHERE delivery_id = $1")
	err = or.DB.DB.QueryRow(query, deliveryId).Scan(&delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region, &delivery.Email)
	if err != nil {
		return domain.Order{}, err
	}

	var payment domain.Payment
	query = fmt.Sprintf("SELECT transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee FROM payment WHERE payment_id = $1")
	err = or.DB.DB.QueryRow(query, paymentId).Scan(&payment.Transaction, &payment.RequestID, &payment.Currency, &payment.Provider, &payment.Amount, &payment.PaymentDT, &payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)
	if err != nil {
		return domain.Order{}, err
	}

	query = fmt.Sprintf("SELECT chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status FROM item WHERE order_id = $1")
	rows, err := or.DB.DB.Query(query, orderId)
	if err != nil {
		return domain.Order{}, err
	}
	defer rows.Close()

	var items []domain.Item
	for rows.Next() {
		item := domain.Item{}
		err := rows.Scan(&item.ChrtID, &item.TrackNumber, &item.Price, &item.RID, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status)
		if err != nil {
			return domain.Order{}, err
		}
		items = append(items, item)
	}

	var order domain.Order
	query = fmt.Sprintf("SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard FROM  \"Order\" WHERE order_uid = $1")
	err = or.DB.DB.QueryRow(query, uid).Scan(&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService, &order.ShardKey, &order.SMID, &order.DateCreated, &order.OOFShard)
	if err != nil {
		return domain.Order{}, err
	}

	order.Items = items
	order.Payment = payment
	order.Delivery = delivery

	return order, nil
}
