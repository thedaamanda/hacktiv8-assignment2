package models

// Items represents the items in an order
type Item struct {
	ItemId      uint   `gorm:"primaryKey" json:"itemId"`
	ItemCode    string `gorm:"not null;type:varchar(50)" json:"itemCode"`
	Description string `gorm:"not null;type:varchar(255)" json:"description"`
	Quantity    int    `gorm:"not null;type:int" json:"quantity"`
	OrderID     uint   `gorm:"not null" json:"orderId"`
}
