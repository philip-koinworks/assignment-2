package params

type CreateOrder struct {
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"ordered_at"`
}

type CreateOrderItem struct {
	CreateOrder
	I []CreateItem
}
