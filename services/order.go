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

func (p *OrderService) CreateOrder(request params.CreateOrder) *params.Response {
	model := models.Order{
		CustomerName: request.CustomerName,
	}

	err := p.orderRepo.CreateOrder(&model)
	if err != nil {
		return &params.Response{
			Status:         400,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		}
	}

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
