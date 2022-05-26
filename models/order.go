package models

type Order struct {
	Model
	CustomerName string `json:"customerName"`
	Items        []Item `gorm:"foreignKey:OrderId;references:ID;constraint:OnDelete:CASCADE"`
}
