package main

import (
	"deliveryCheck/pkg/natsutil"
	"github.com/nats-io/nats.go"
	"log"
)

func main() {
	// Подключение к серверу NATS
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Подписка на канал
	_, err = nc.Subscribe("subject", natsutil.MessageHandler)
	if err != nil {
		log.Fatal(err)
	}

	// Тестовая отправка сообщения
	err = nc.Publish("subject", []byte("Hello!"))
	if err != nil {
		log.Fatal(err)
	}
}
