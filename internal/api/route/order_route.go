package route

import (
	"deliveryCheck/internal/lru"
	"github.com/gin-gonic/gin"
	"time"

	"deliveryCheck/app"
	"deliveryCheck/internal/api/controller"
	"deliveryCheck/internal/repository"
	"deliveryCheck/internal/usecase"
)

func NewOrderRoute(env *app.Env, db app.Postgres, group *gin.RouterGroup) {
	or := repository.NewOrderRepository(db)
	oc := &controller.OrderController{
		OrderCache:   lru.New(30 * time.Second),
		OrderUsecase: usecase.NewOrderUsecase(or),
		Env:          env,
	}

	group.GET("/order/:orderuid", oc.GetOrderByUID)
}
