package controllers

import (
	"assignment-2/params"
	"assignment-2/services"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

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
	var reqOrderItem params.CreateOrderItem

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		})
		return
	}

	err = json.Unmarshal(body, &reqOrderItem)

	response := p.orderService.CreateOrder(reqOrderItem)
	c.JSON(response.Status, response)
}

func (p *OrderController) GetAllOrders(c *gin.Context) {
	response := p.orderService.GetAllOrders()
	c.JSON(response.Status, response)
}

func (p *OrderController) DeleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err,
		})
		return
	}
	response := p.orderService.DeleteOrder(id)
	c.JSON(response.Status, response)
}
