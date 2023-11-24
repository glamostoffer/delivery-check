package route

import (
	"github.com/gin-gonic/gin"

	"deliveryCheck/app"
)

func Setup(env *app.Env, db app.Postgres, gin *gin.Engine) {
	publicRouter := gin.Group("")

	NewOrderRoute(env, db, publicRouter)
}
