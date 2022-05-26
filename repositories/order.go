package repositories

import (
	"assignment-2/models"

	"gorm.io/gorm"
)

type OrderRepo interface {
	CreateOrder(order *models.Order) error
	GetAllOrders() (*[]models.Order, error)
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *orderRepo {
	return &orderRepo{db}
}

func (p *orderRepo) CreateOrder(order *models.Order) error {
	return p.db.Create(order).Error
}

func (p *orderRepo) GetAllOrders() (*[]models.Order, error) {
	var person []models.Order
	err := p.db.Find(&person).Error
	return &person, err
}
