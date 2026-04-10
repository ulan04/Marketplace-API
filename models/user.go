package models

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Phone string `json:"phone" gorm:"not null"`
	Items []Item `json:"items,omitempty"`
}