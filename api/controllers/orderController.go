package controllers

import (
	"assignment2/api/requests"
	"assignment2/api/responses"
	"assignment2/api/services"
	"assignment2/models"
	"assignment2/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary Create a new order
// @Description Create a new order and save to database
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body requests.CreateOrder true "Create Order Request"
// @Success 201 {object} models.Order
// @Failure 400 {object} responses.ErrorValidation
// @Failure 500 {object} responses.Error
// @Router /orders [post]
func CreateOrder(ctx *gin.Context) {
	var (
		orderPayload requests.CreateOrder
		order        models.Order
		err          error
	)

	if err = ctx.ShouldBind(&orderPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorValidation{
			Message: "Invalid request body",
			Errors:  utils.ParseError(err),
		})
		return
	}

	orderPayload.BindToModelOrder(&order)

	if err = services.Create(&order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Error{
			Message: "Error creating order",
			Error:   err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, responses.Data{
		Data:    order,
		Message: "Order created successfully",
	})
}

// GetOrders godoc
// @Summary Get all orders
// @Description Get all orders from database
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {object} responses.DataList
// @Failure 500 {object} responses.Error
// @Router /orders [get]
func GetOrders(ctx *gin.Context) {
	var (
		orders []models.Order
		err    error
	)

	if orders, err = services.GetAll(); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Error{
			Message: "Error getting orders",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.DataList{
		Count:   len(orders),
		Data:    orders,
		Message: "Orders retrieved successfully",
	})
}

// GetOrdersByID godoc
// @Summary Get orders by ID
// @Description Get all orders by user ID from database
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} responses.Data
// @Failure 404 {object} responses.Error
// @Router /orders/{id} [get]
func GetOrderByID(ctx *gin.Context) {
	var (
		id    = ctx.Param("id")
		order models.Order
		err   error
	)

	if err = services.GetByID(&order, id); err != nil {
		ctx.JSON(http.StatusNotFound, responses.Error{
			Message: "Order ID not found",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Data{
		Data:    order,
		Message: "Order ID retrieved successfully",
	})
}

// Update godoc
// @Summary Update an order
// @Description Update an order by ID and save to database
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Param order body requests.UpdateOrder true "Update Order Request"
// @Success 200 {object} models.Order
// @Failure 400 {object} responses.ErrorValidation
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /orders/{id} [put]
func UpdateOrder(ctx *gin.Context) {
	var (
		id           = ctx.Param("id")
		orderPayload requests.UpdateOrder
		order        models.Order
		orderUpdate  models.Order
		err          error
	)

	if err = services.GetByID(&orderUpdate, id); err != nil {
		ctx.JSON(http.StatusNotFound, responses.Error{
			Message: "Order ID not found",
			Error:   err.Error(),
		})
		return
	}

	if err = ctx.ShouldBind(&orderPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorValidation{
			Message: "Invalid request body",
			Errors:  utils.ParseError(err),
		})
		return
	}

	orderPayload.BindToModelOrderUpdate(&orderUpdate)

	if err = services.Update(&order, &orderUpdate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Error{
			Message: "Error updating order",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.Data{
		Data:    order,
		Message: "Order updated successfully",
	})
}

// Delete godoc
// @Summary Delete an order
// @Description Delete an order by ID from database
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} responses.DataAffected
// @Failure 404 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /orders/{id} [delete]
func DeleteOrder(ctx *gin.Context) {
	var (
		id      = ctx.Param("id")
		deleted int64
		order   models.Order
		err     error
	)

	if err = services.GetByID(&order, id); err != nil {
		ctx.JSON(http.StatusNotFound, responses.Error{
			Message: "Order ID not found",
			Error:   err.Error(),
		})
		return
	}

	if deleted, err = services.Delete(&order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, responses.Error{
			Message: "Error deleting order",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responses.DataAffected{
		Count:   int(deleted),
		Message: "Order deleted successfully",
	})
}
