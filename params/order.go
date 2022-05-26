package params

type CreateOrder struct {
	CustomerName string `json:"customerName"`
}

type CreateOrderItem struct {
	CreateOrder
	I []CreateItem
}
