package controller

import (
	"deliveryCheck/app"
	"deliveryCheck/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController struct {
	OrderUsecase domain.OrderUsecase
	Env          *app.Env
}

type OrderParams struct {
	UID string `url:"uid"`
}

func (oc OrderController) GetOrderByUID(c *gin.Context) {
	var params OrderParams
	err := c.ShouldBindUri(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	order, err := oc.OrderUsecase.GetOrderByUID(c, params.UID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}
