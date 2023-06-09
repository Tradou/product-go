package handlers

import (
	"github.com/gin-gonic/gin"
	"main/infrastructure/database"
	"main/infrastructure/models"
	"net/http"
)

func GetOrders(c *gin.Context) {
	var orders []models.GetOrder

	db, err := database.GetDBConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to connect database")
		return
	}

	if err := db.
		Preload("OrderItems").
		//Preload("OrderItems", func(db *gorm.DB) *gorm.DB {
		//	return db.Where("order_items.attribute_product_id IS NULL").
		//		Preload("Product")
		//}).
		//Preload("OrderItems", func(db *gorm.DB) *gorm.DB {
		//	return db.Where("order_items.attribute_product_id IS NOT NULL").
		//		Preload("AttributeProduct.Product").
		//		Preload("AttributeProduct.Attribute")
		//}).
		Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	var order models.Order
	orderId := c.Param("id")

	db, err := database.GetDBConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to connect database")
		return
	}

	if err := db.First(&order, orderId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to retrieve data")
		return
	}

	c.JSON(http.StatusOK, order)
}
