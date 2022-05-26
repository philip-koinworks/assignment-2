package models

type Order struct {
	Model
	CustomerName string `json:"customerName"`
	Items        []Item `gorm:"foreignKey:OrderId;references:ID"`
}

type OrderItem struct {
	Order
	Items []Item
}
