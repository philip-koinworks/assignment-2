package models

type Item struct {
	Model
	OrderId         uint   `json:"orderId" gorm:"foreignKey:ID"`
	ItemCode        int    `json:"itemCode"`
	ItemDescription string `json:"itemDescription"`
	ItemQuantity    int    `json:"itemQuantity"`
}
