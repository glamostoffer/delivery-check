package route

import (
	"github.com/gin-gonic/gin"

	"deliveryCheck/app"
	"deliveryCheck/internal/api/controller"
	"deliveryCheck/internal/repository"
	"deliveryCheck/internal/usecase"
)

func NewOrderRoute(env *app.Env, db app.Postgres, group *gin.RouterGroup) {
	or := repository.NewOrderRepository(db)
	oc := &controller.OrderController{
		OrderUsecase: usecase.NewOrderUsecase(or),
		Env:          env,
	}

	group.GET("/order/:orderuid", oc.GetOrderByUID)
}
