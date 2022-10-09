package models

import "time"

// Order represents the order of a customer and the items in the order
type Order struct {
	OrderId      uint      `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `gorm:"not null;type:varchar(150)" json:"customerName"`
	OrderedAt    time.Time `gorm:"autoCreateTime" json:"orderedAt"`
	Items        []Item    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"items"`
}
