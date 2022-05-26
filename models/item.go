package models

type Item struct {
	Model
	OrderId         uint   `json:"order_id"`
	ItemCode        int    `json:"item_code"`
	ItemDescription string `json:"item_description"`
	ItemQuantity    int    `json:"item_quantity"`
}
