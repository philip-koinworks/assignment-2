package repositories

import (
	"assignment-2/models"

	"gorm.io/gorm"
)

type OrderRepo interface {
	CreateOrder(order *models.Order) (uint, error)
	GetAllOrders() (*[]models.Order, error)
	CreateItemOrder(item *[]models.Item) error
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *orderRepo {
	return &orderRepo{db}
}

func (p *orderRepo) CreateOrder(order *models.Order) (uint, error) {
	res := p.db.Create(&order)
	if res.Error != nil {
		return 0, res.Error
	}
	return order.ID, nil
}

func (p *orderRepo) CreateItemOrder(item *[]models.Item) error {
	res := p.db.Create(&item)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (p *orderRepo) GetAllOrders() (*[]models.Order, error) {
	var person []models.Order
	err := p.db.Find(&person).Error
	return &person, err
}
