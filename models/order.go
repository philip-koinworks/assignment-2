package models

type Order struct {
	Model
	CustomerName string `json:"customerName"`
	OrderedAt    string `json:"orderedAt"`
}

type OrderItem struct {
	Order
	I []Item
}
