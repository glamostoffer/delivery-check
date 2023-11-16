package main

import (
	"deliveryCheck/app"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	app := app.App()
	defer app.CloseDBConnection()

	env := app.Env

	gin := gin.Default()

	//route.Setup(env, timeout, db, gin)

	gin.Run(fmt.Sprintf("%s:%s", env.ServerHost, env.ServerPort))
}
