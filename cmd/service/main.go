package main

import (
	"deliveryCheck/app"
	br "deliveryCheck/internal/broker"
	"deliveryCheck/internal/repository"
	"deliveryCheck/internal/usecase"
	"deliveryCheck/pkg/natsconnector"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func startServer(app app.Application) {
	env := app.Env

	gin := gin.Default()

	//route.Setup(env, timeout, db, gin)

	gin.Run(fmt.Sprintf("%s:%s", env.ServerHost, env.ServerPort))
}

func connectNats(app app.Application) {
	env := app.Env

	// Подключение к серверу NATS
	nc, err := natsconnector.NewNatsConnector(env)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Client.Close()

	// Подписка на канал
	broker := br.NewNatsBroker(
		nc,
		usecase.NewNatsUsecase(
			repository.NewOrderRepository(app.Postgres),
		),
	)

	err = broker.Subscribe(app.Env.ChannelName)
	if err != nil {
		log.Fatal(err)
	}

	// Тестовая отправка сообщения
	//err = nc.Publish(env.ChannelName, []byte("Hello!"))
	//if err != nil {
	//	log.Fatal(err)
	//}
}

func main() {
	app := app.App()
	defer app.CloseDBConnection()

	go connectNats(app)
	startServer(app)
}
