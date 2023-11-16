package main

import (
	"deliveryCheck/app"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"log"
)

func MessageHandler(m *nats.Msg) {
	log.Printf("Received a message: %s", string(m.Data))
}

func main() {
	app := app.App()
	defer app.CloseDBConnection()

	env := app.Env

	gin := gin.Default()

	//route.Setup(env, timeout, db, gin)

	gin.Run(fmt.Sprintf("%s:%s", env.ServerHost, env.ServerPort))
}

//func main() {
//	// Подключение к серверу NATS
//	nc, err := nats.Connect("nats://localhost:4222")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer nc.Close()
//
//	// Подписка на канал
//	_, err = nc.Subscribe("subject", natsutil.MessageHandler)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Тестовая отправка сообщения
//	err = nc.Publish("subject", []byte("Hello!"))
//	if err != nil {
//		log.Fatal(err)
//	}
//}
