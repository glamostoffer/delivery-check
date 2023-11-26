package controller

import (
	"deliveryCheck/internal/lru"
	"deliveryCheck/internal/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"deliveryCheck/app"
	"deliveryCheck/internal/domain"
)

type OrderController struct {
	OrderCache   *lru.Cache
	OrderUsecase usecase.OrderUsecase
	Env          *app.Env
}

type OrderParams struct {
	UID string `url:"uid"`
}

func (oc OrderController) GetOrderByUID(c *gin.Context) {
	uid := c.Param("orderuid")

	order, found := oc.OrderCache.Get(uid)
	if !found {
		order, err := oc.OrderUsecase.GetOrderByUID(c, uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}
		oc.OrderCache.Set(uid, order, 20*time.Second)
		c.JSON(http.StatusOK, order)
		return
	}

	c.JSON(http.StatusOK, order)
}
