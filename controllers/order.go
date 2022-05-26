package controllers

import (
	"assignment-2/params"
	"assignment-2/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{
		orderService: *service,
	}
}

func (p *OrderController) CreateNewOrder(c *gin.Context) {
	var req params.CreateOrder

	err := c.ShouldBindJSON(&req)
	fmt.Println(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		})
		return
	}

	response := p.orderService.CreateOrder(req)
	c.JSON(response.Status, response)
}

func (p *OrderController) GetAllOrders(c *gin.Context) {
	response := p.orderService.GetAllOrders()
	c.JSON(response.Status, response)
}
