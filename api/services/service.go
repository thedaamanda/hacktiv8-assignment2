package services

import (
	"assignment2/config"
	"assignment2/models"
)

func Create(newOrder *models.Order) error {
	err := config.GetDB().Create(newOrder).Error
	return err
}

func GetAll() (orders []models.Order, err error) {
	err = config.GetDB().Preload("Items").Find(&orders).Error
	return
}

func GetByID(order *models.Order, id string) error {
	err := config.GetDB().Preload("Items").First(&order, id).Error
	return err
}

func Update(order, orderUpdate *models.Order) error {
	order.OrderId = orderUpdate.OrderId
	order.CustomerName = orderUpdate.CustomerName
	order.OrderedAt = orderUpdate.OrderedAt
	order.Items = orderUpdate.Items

	config.GetDB().Delete(&orderUpdate.Items, "order_id = ?", orderUpdate.OrderId)

	err := config.GetDB().Save(&order).Error
	return err
}

func Delete(order *models.Order) (int64, error) {
	result := config.GetDB().Delete(&order)
	return result.RowsAffected, result.Error
}
