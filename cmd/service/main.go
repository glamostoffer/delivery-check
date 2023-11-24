package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"deliveryCheck/app"
	"deliveryCheck/internal/api/route"
	br "deliveryCheck/internal/broker"
	"deliveryCheck/internal/repository"
	"deliveryCheck/internal/usecase"
	"deliveryCheck/pkg/generator"
	"deliveryCheck/pkg/natsconnector"
)

func startServer(app app.Application) {
	env := app.Env

	gin := gin.Default()

	gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:63342"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Access-Control-Allow-Origin"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:63342"
		},
		MaxAge: 12 * time.Hour,
	}))

	route.Setup(env, app.Postgres, gin)

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

	// Тестовая отправка сообщений
	broker2 := br.NewNatsBroker(
		nc,
		usecase.NewNatsUsecase(
			repository.NewOrderRepository(app.Postgres),
		),
	)

	err = broker2.Subscribe(app.Env.ChannelName)
	if err != nil {
		log.Fatal(err)
	}

	err = broker2.Publish(env.ChannelName, "Hello!")
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := generator.GenerateRandomJSON()
		if err != nil {
			log.Printf("error sending message: %s", err)
		}
		err = broker2.Publish(env.ChannelName, msg)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(10 * time.Second)
	}
}

func main() {
	app := app.App()
	defer app.CloseDBConnection()

	go connectNats(app)
	startServer(app)
}
