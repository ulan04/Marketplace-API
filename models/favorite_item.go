package models

import "time"

type FavoriteItem struct {
	UserID    uint      `json:"user_id" gorm:"primaryKey"`
	ItemID    uint      `json:"item_id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`

	User User `json:"user,omitempty" gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Item Item `json:"item,omitempty" gorm:"foreignKey:ItemID;references:ID;constraint:OnDelete:CASCADE"`
}

func (FavoriteItem) TableName() string {
	return "favorite_items"
}