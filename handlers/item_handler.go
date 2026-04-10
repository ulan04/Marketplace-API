package handlers

import (
	"net/http"
	"strconv"

	"marketplace-api/config"
	"marketplace-api/models"

	"github.com/gin-gonic/gin"
)

func GetItems(c *gin.Context) {
	var items []models.Item

	if err := config.DB.Preload("User").Preload("Category").Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}

	c.JSON(http.StatusOK, items)
}

func CreateItem(c *gin.Context) {
	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	if item.Title == "" || item.Price <= 0 || item.Location == "" || item.UserID == 0 || item.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title, price, location, user_id and category_id are required",
		})
		return
	}

	var user models.User
	if err := config.DB.First(&user, item.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var category models.Category
	if err := config.DB.First(&category, item.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		return
	}

	if err := config.DB.Create(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create item"})
		return
	}

	config.DB.Preload("User").Preload("Category").First(&item, item.ID)

	c.JSON(http.StatusCreated, item)
}

func GetItemByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var item models.Item
	if err := config.DB.Preload("User").Preload("Category").First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, item)
}

func UpdateItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	var updatedItem models.Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	if updatedItem.Title == "" || updatedItem.Price <= 0 || updatedItem.Location == "" || updatedItem.UserID == 0 || updatedItem.CategoryID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title, price, location, user_id and category_id are required",
		})
		return
	}

	var user models.User
	if err := config.DB.First(&user, updatedItem.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	var category models.Category
	if err := config.DB.First(&category, updatedItem.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		return
	}

	item.Title = updatedItem.Title
	item.Description = updatedItem.Description
	item.Price = updatedItem.Price
	item.Location = updatedItem.Location
	item.UserID = updatedItem.UserID
	item.CategoryID = updatedItem.CategoryID

	if err := config.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item"})
		return
	}

	config.DB.Preload("User").Preload("Category").First(&item, item.ID)

	c.JSON(http.StatusOK, item)
}

func DeleteItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	if err := config.DB.Delete(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted successfully"})
}