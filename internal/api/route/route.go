package route

import (
	"deliveryCheck/app"
	"github.com/gin-gonic/gin"
)

func Setup(env *app.Env, db *app.Postgres, gin *gin.Engine) {
	//publicRouter := gin.Group("")

	//NewMessageRouter(env, db, publicRouter)
}
