package requests

import (
	. "assignment2/models"
	"time"
)

type UpdateOrder struct {
	CustomerName string       `json:"customerName" binding:"required"`
	OrderedAt    time.Time    `json:"orderedAt" binding:"required"`
	Items        []UpdateItem `json:"items" binding:"required"`
}

type UpdateItem struct {
	ItemCode    string `json:"itemCode" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required"`
}

func (from UpdateOrder) BindToModelOrderUpdate(to *Order) {
	to.CustomerName = from.CustomerName
	to.OrderedAt = from.OrderedAt
	to.Items = ItemOnUpdateMapper(from.Items)
}

func (from UpdateItem) BindToModelItemUpdate(to *Item) {
	to.ItemCode = from.ItemCode
	to.Description = from.Description
	to.Quantity = int(from.Quantity)
}

func ItemOnUpdateMapper(from []UpdateItem) []Item {
	var to []Item
	for _, item := range from {
		var newItem Item
		item.BindToModelItemUpdate(&newItem)
		to = append(to, newItem)
	}
	return to
}
