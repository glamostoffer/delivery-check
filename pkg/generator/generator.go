package generator

import (
	"deliveryCheck/internal/domain"
	"encoding/json"
	"math/rand"
	"time"
)

func generateTimestampString() string {
	t := time.Now().UTC()

	timestampString := t.Format("2006-01-02T15:04:05Z")

	return timestampString
}

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomJSON() (string, error) {
	delivery := domain.Delivery{
		Name:    generateRandomString(10),
		Phone:   generateRandomString(10),
		Zip:     generateRandomString(10),
		City:    generateRandomString(10),
		Address: generateRandomString(10),
		Region:  generateRandomString(10),
		Email:   generateRandomString(10),
	}

	payment := domain.Payment{
		Transaction:  generateRandomString(10),
		RequestID:    generateRandomString(10),
		Currency:     generateRandomString(3),
		Provider:     generateRandomString(10),
		Amount:       rand.Int() % 10000,
		PaymentDT:    rand.Int63() % 2147483646,
		Bank:         generateRandomString(10),
		DeliveryCost: rand.Int() % 2147483646,
		GoodsTotal:   rand.Int() % 2147483646,
		CustomFee:    rand.Int() % 2147483646,
	}

	order := domain.Order{
		OrderUID:          payment.Transaction,
		TrackNumber:       generateRandomString(10),
		Entry:             generateRandomString(4),
		Delivery:          delivery,
		Payment:           payment,
		Locale:            generateRandomString(2),
		InternalSignature: generateRandomString(10),
		CustomerID:        generateRandomString(10),
		DeliveryService:   generateRandomString(10),
		ShardKey:          generateRandomString(10),
		SMID:              rand.Int() % 2147483646,
		DateCreated:       generateTimestampString(),
		OOFShard:          generateRandomString(10),
	}

	var items []domain.Item
	for i := 0; i < rand.Int()%50; i += 1 {
		item := domain.Item{
			ChrtID:      rand.Int() % 2147483646,
			TrackNumber: order.TrackNumber,
			Price:       rand.Int() % 2147483646,
			RID:         generateRandomString(12),
			Name:        generateRandomString(10),
			Sale:        rand.Int() % 2147483646,
			Size:        generateRandomString(10),
			TotalPrice:  rand.Int() % 2147483646,
			NmID:        rand.Int() % 2147483646,
			Brand:       generateRandomString(10),
			Status:      202,
		}

		items = append(items, item)
	}

	order.Items = items

	// Кодируем экземпляр структуры в JSON-строку
	jsonData, err := json.Marshal(order)
	if err != nil {
		return "", err
	}

	// Возвращаем JSON-строку
	return string(jsonData), nil
}
