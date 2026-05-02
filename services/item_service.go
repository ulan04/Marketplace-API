package services

import (
	"errors"
	"marketplace-api/config"
	"marketplace-api/models"
)

type ItemService struct{}

func NewItemService() *ItemService {
	return &ItemService{}
}

func (s *ItemService) GetItems() ([]models.Item, error) {
	var items []models.Item
	err := config.DB.Preload("User").Preload("Category").Find(&items).Error
	return items, err
}

func (s *ItemService) GetItemByID(id uint) (models.Item, error) {
	var item models.Item
	err := config.DB.Preload("User").Preload("Category").First(&item, id).Error
	return item, err
}

func (s *ItemService) CreateItem(item *models.Item) error {
	if item.Title == "" || item.Price <= 0 || item.Location == "" || item.UserID == 0 || item.CategoryID == 0 {
		return errors.New("invalid item data")
	}

	var user models.User
	if err := config.DB.First(&user, item.UserID).Error; err != nil {
		return errors.New("user not found")
	}

	var category models.Category
	if err := config.DB.First(&category, item.CategoryID).Error; err != nil {
		return errors.New("category not found")
	}

	if err := config.DB.Create(item).Error; err != nil {
		return err
	}

	return config.DB.Preload("User").Preload("Category").First(item, item.ID).Error
}

func (s *ItemService) UpdateItem(id uint, updatedItem *models.Item) error {
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return err
	}

	if updatedItem.Title == "" || updatedItem.Price <= 0 || updatedItem.Location == "" || updatedItem.UserID == 0 || updatedItem.CategoryID == 0 {
		return errors.New("invalid item data")
	}

	var user models.User
	if err := config.DB.First(&user, updatedItem.UserID).Error; err != nil {
		return errors.New("user not found")
	}

	var category models.Category
	if err := config.DB.First(&category, updatedItem.CategoryID).Error; err != nil {
		return errors.New("category not found")
	}

	return config.DB.Model(&item).Updates(updatedItem).Error
}

func (s *ItemService) DeleteItem(id uint) error {
	return config.DB.Delete(&models.Item{}, id).Error
}