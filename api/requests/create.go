package requests

import (
	. "assignment2/models"
	"time"
)

type CreateOrder struct {
	CustomerName string       `json:"customerName" binding:"required"`
	OrderedAt    time.Time    `json:"orderedAt" binding:"required"`
	Items        []CreateItem `json:"items" binding:"required"`
}

type CreateItem struct {
	ItemCode    string `json:"itemCode" binding:"required"`
	Description string `json:"description" binding:"required"`
	Quantity    uint   `json:"quantity" binding:"required"`
}

func (from CreateOrder) BindToModelOrder(to *Order) {
	to.CustomerName = from.CustomerName
	to.OrderedAt = from.OrderedAt
	to.Items = ItemOnCreateMapper(from.Items)
}

func (from CreateItem) BindToModelItem(to *Item) {
	to.ItemCode = from.ItemCode
	to.Description = from.Description
	to.Quantity = int(from.Quantity)
}

func ItemOnCreateMapper(from []CreateItem) []Item {
	var to []Item
	for _, item := range from {
		var newItem Item
		item.BindToModelItem(&newItem)
		to = append(to, newItem)
	}
	return to
}
