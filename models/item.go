package models

import "time"

type Item struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Price       float64   `json:"price" gorm:"not null"`
	Location    string    `json:"location" gorm:"not null"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	CategoryID  uint      `json:"category_id" gorm:"not null"`
	User        User      `json:"user"`
	Category    Category  `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}