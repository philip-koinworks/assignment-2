package params

type CreateItem struct {
	OrderId         int    `json:"order_id"`
	ItemCode        string `json:"item_code"`
	ItemDescription string `json:"item_description"`
	ItemQuantity    string `json:"item_quantity"`
}
