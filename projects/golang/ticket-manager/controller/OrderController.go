package controller

import (
	"net/http"
	"strconv"
	"ticket-manager/model"
	"ticket-manager/service"

	tracing "github.com/codeandcode0x/traceandtrace-go"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	apiVersion string
	Service    *service.OrderService
}

// get controller
func (uc *OrderController) getOrderController() *OrderController {
	var svc *service.OrderService
	return &OrderController{"v1", svc}
}

// create user
func (uc *OrderController) CreateOrder(c *gin.Context) {
	name := c.PostForm("name")
	money, exists := c.GetPostForm("money")
	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": "money is null",
		})
	}

	moneyUint64, _ := strconv.ParseUint(money, 10, 64)

	order := model.Order{
		Name:  name,
		Money: moneyUint64,
	}

	err := uc.getOrderController().Service.CreateOrder(&order)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": order,
	})
}

func (uc *OrderController) GetAllOrders(c *gin.Context) {
	orders, err := uc.getOrderController().Service.FindAllOrders()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": orders,
	})
}

func (uc *OrderController) GetOrderByPages(c *gin.Context) {
	// add tracing
	_, cancel := tracing.AddHttpTracing("OrderService", c.Request.Header, map[string]string{})
	defer cancel()

	orders, err := uc.getOrderController().Service.FindOrderByPages(1, 1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -1,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": orders,
	})
}
