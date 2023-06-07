package handlers

import (
	"github.com/gin-gonic/gin"
	"main/infrastructure/database"
	"main/infrastructure/models"
	"net/http"
)

func GetProducts(c *gin.Context) {
	var products []models.Product

	db, err := database.GetDBConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to connect database")
		return
	}

	if err := db.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to retrieve data")
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	var product models.Product
	productId := c.Param("id")

	db, err := database.GetDBConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to connect database")
		return
	}

	if err := db.First(&product, productId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to retrieve data")
		return
	}

	c.JSON(http.StatusOK, product)
}

func StoreProduct(c *gin.Context) {
	var product models.StoreProduct

	if err := c.ShouldBind(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.GetDBConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to connect database")
		return
	}

	if err := db.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	productID := c.Param("id")
	var updateData models.UpdateProduct

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.GetDBConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	var existingProduct models.Product
	if err := db.First(&existingProduct, productID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data from the database"})
		return
	}

	if err := db.Model(&existingProduct).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, existingProduct)
}
