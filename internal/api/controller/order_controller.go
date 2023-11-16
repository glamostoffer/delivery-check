package controller

import (
	"deliveryCheck/app"
	"deliveryCheck/internal/domain"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderUsecase domain.OrderUsecase
	Env          *app.Env
}

func (oc OrderController) GetOrderByUID(c *gin.Context) {

}
