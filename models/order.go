package models

type Order struct {
	Model
	CustomerName string `json:"customerName"`
}

type OrderItem struct {
	Order
	Items []Item
}
