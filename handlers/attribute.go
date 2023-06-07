package handlers

import (
	"github.com/gin-gonic/gin"
	"main/infrastructure/database"
	"main/infrastructure/models"
	"net/http"
)

func GetAttributes(c *gin.Context) {
	var attributes []models.Attribute

	db, err := database.GetDBConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to connect database")
		return
	}

	if err := db.Find(&attributes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, attributes)
}

func GetAttribute(c *gin.Context) {
	var attribute models.Attribute
	attributeId := c.Param("id")

	db, err := database.GetDBConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to connect database")
		return
	}

	if err := db.First(&attribute, attributeId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to retrieve data")
		return
	}

	c.JSON(http.StatusOK, attribute)
}

func StoreAttribute(c *gin.Context) {
	var attribute models.StoreAttribute

	if err := c.ShouldBind(&attribute); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.GetDBConnection()

	if err != nil {
		c.JSON(http.StatusInternalServerError, "Failed to connect database")
		return
	}

	if err := db.Create(&attribute).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, attribute)
}

func UpdateAttribute(c *gin.Context) {
	attributeId := c.Param("id")
	var updateData models.UpdateAttribute

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.GetDBConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	var attribute models.Attribute
	if err := db.First(&attribute, attributeId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve data from the database"})
		return
	}

	if err := db.Model(&attribute).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update attribute"})
		return
	}

	c.JSON(http.StatusOK, attribute)
}
