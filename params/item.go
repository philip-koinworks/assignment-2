package params

type CreateItem struct {
	OrderId         uint   `json:"orderId"`
	ItemCode        int    `json:"itemCode"`
	ItemDescription string `json:"itemDescription"`
	ItemQuantity    int    `json:"itemQuantity"`
}
