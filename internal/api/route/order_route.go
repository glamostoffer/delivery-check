package route

import (
	"deliveryCheck/app"
	"deliveryCheck/internal/api/controller"
	"deliveryCheck/internal/repository"
	"deliveryCheck/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewOrderRoute(env *app.Env, db app.Postgres, group *gin.RouterGroup) {
	or := repository.NewOrderRepository(db)
	oc := &controller.OrderController{
		OrderUsecase: usecase.NewOrderUsecase(or),
		Env:          env,
	}

	group.GET("/order", oc.GetOrderByUID)
}
