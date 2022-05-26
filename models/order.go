package models

type Order struct {
	Model
	CustomerName string `json:"customerName"`
}

type OrderItem struct {
	Order
	I []Item
}
