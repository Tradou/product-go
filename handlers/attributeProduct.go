package handlers

import (
	"github.com/gin-gonic/gin"
	"main/infrastructure/database"
	"main/infrastructure/models"
	"net/http"
)

func GetAttributeProducts(c *gin.Context) {
	var attributeProducts []models.AttributeProduct

	db, err := database.GetDBConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to connect database")
		return
	}

	if err := db.Find(&attributeProducts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, attributeProducts)
}

func GetAttributeProduct(c *gin.Context) {
	var attributeProduct models.AttributeProduct
	attributeProductId := c.Param("id")

	db, err := database.GetDBConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to connect database")
		return
	}

	if err := db.First(&attributeProduct, attributeProductId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to retrieve data")
		return
	}

	c.JSON(http.StatusOK, attributeProduct)
}

func StoreAttributeProduct(c *gin.Context) {
	var attributeProduct models.StoreAttributeProduct

	if err := c.ShouldBind(&attributeProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.GetDBConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to connect database")
		return
	}

	if err := db.Create(&attributeProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, attributeProduct)
}

func UpdateAttributeProduct(c *gin.Context) {
	attributeProductId := c.Param("id")
	var updateData models.UpdateAttributeProduct

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.GetDBConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	var attributeProduct models.AttributeProduct
	if err := db.First(&attributeProduct, attributeProductId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data from the database"})
		return
	}

	if err := db.Model(&attributeProduct).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update attribute product"})
		return
	}

	c.JSON(http.StatusOK, attributeProduct)
}
