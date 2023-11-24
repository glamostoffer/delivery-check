package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"deliveryCheck/app"
	"deliveryCheck/internal/domain"
)

type OrderController struct {
	OrderUsecase domain.OrderUsecase
	Env          *app.Env
}

type OrderParams struct {
	UID string `url:"uid"`
}

func (oc OrderController) GetOrderByUID(c *gin.Context) {
	uid := c.Param("orderuid")

	order, err := oc.OrderUsecase.GetOrderByUID(c, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
