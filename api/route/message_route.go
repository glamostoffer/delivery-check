package route

import (
	"deliveryCheck/app"
	"github.com/gin-gonic/gin"
)

func NewMessageRouter(env *app.Env, db *app.Postgres, group *gin.RouterGroup) {
	//mr := repository.NewMessageRepository(db)
	//mc := &controller.MessageController{
	//	MessageUsecase: usecase.NewMessageUsecase(ur),
	//	Env:            env,
	//}
	//group.POST("/", mc.Method)
}
