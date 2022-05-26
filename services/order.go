package services

import (
	"assignment-2/models"
	"assignment-2/params"
	"assignment-2/repositories"
	"net/http"

	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo repositories.OrderRepo
}

var db = gorm.DB{}
var repo = repositories.NewOrderRepo(&db)

func NewOrderService(repo repositories.OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: repo,
	}
}

func (p *OrderService) CreateOrder(request params.CreateOrderItem) *params.Response {
	orderModel := models.Order{CustomerName: request.CustomerName}
	orderItem := []models.Item{}

	id, err := p.orderRepo.CreateOrder(&orderModel)

	if err != nil {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

	for _, item := range request.Items {
		modelItem := models.Item{
			OrderId:         id,
			ItemCode:        item.ItemCode,
			ItemDescription: item.ItemDescription,
			ItemQuantity:    item.ItemQuantity,
		}
		orderItem = append(orderItem, modelItem)
	}

	err = p.orderRepo.CreateItemOrder(&orderItem)

	return &params.Response{
		Status:  200,
		Message: "CREATE SUCCESS",
		Payload: request,
	}
}

func (p *OrderService) GetAllOrders() *params.Response {
	response, err := p.orderRepo.GetAllOrders()
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: response,
	}
}

func (p *OrderService) DeleteOrder(id int) *params.Response {
	err := p.orderRepo.DeleteOrder(id)
	if err != nil {
		return &params.Response{
			Status:         http.StatusInternalServerError,
			Error:          "INTERNAL SERVER ERROR",
			AdditionalInfo: err.Error(),
		}
	}

	return &params.Response{
		Status: http.StatusOK,
	}
}
