package models

type Order struct {
	Model
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"ordered_at"`
}

type OrderItem struct {
	Order
	I []Item
}
