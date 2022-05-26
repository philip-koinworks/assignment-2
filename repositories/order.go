package repositories

import (
	"assignment-2/models"
	"assignment-2/params"

	"gorm.io/gorm"
)

type OrderRepo interface {
	CreateOrder(order *models.Order) (uint, error)
	GetAllOrders() (*[]models.Order, error)
	CreateItemOrder(item *[]models.Item) error
	DeleteOrder(id int) error
	UpdateOrder(id int, customerName string) error
	UpdateItem(id int, items []params.CreateItem) error
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
	var order []models.Order
	err := p.db.Preload("Items").Find(&order).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (p *orderRepo) DeleteOrder(id int) error {
	err := p.db.Delete(&models.Order{}, id)
	if err != nil {
		return err.Error
	}

	return nil
}

func (p *orderRepo) UpdateOrder(id int, customerName string) error {
	err := p.db.Model(&models.Order{}).Where("ID = ?", id).Update("CustomerName", customerName).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *orderRepo) UpdateItem(id int, items []params.CreateItem) error {
	err := p.db.Model(&models.Item{}).Where("ID = ?", id).Updates(models.Item{
		ItemCode:        items[0].ItemCode,
		ItemDescription: items[0].ItemDescription,
		ItemQuantity:    items[0].ItemQuantity,
	}).Error
	if err != nil {
		return err
	}

	return nil
}
