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
	orderModel := models.OrderItem{
		Order: models.Order{CustomerName: request.CustomerName},
		Items: []models.Item{},
	}

	id, err := p.orderRepo.CreateOrder(&orderModel.Order)

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
		orderModel.Items = append(orderModel.Items, modelItem)
	}

	err = p.orderRepo.CreateItemOrder(&orderModel.Items)

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
