package repositories

import (
	"assignment-2/models"

	"gorm.io/gorm"
)

type ItemRepo interface {
	CreateItem(order *models.Order) (uint, error)
	GetAllItem() (*[]models.Order, error)
}

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) *itemRepo {
	return &itemRepo{db}
}

func (p *itemRepo) CreateItem(order *models.Item) (uint, error) {
	res := p.db.Create(&order)
	if res.Error != nil {
		return 0, res.Error
	}
	return order.ID, nil
}

func (p *itemRepo) GetAllItem() (*[]models.Item, error) {
	var person []models.Item
	err := p.db.Find(&person).Error
	return &person, err
}
