package models

import (
	"gorm.io/gorm"
)

// WishList -> Model for Wishlist Table
type WishList struct {
	gorm.Model
	ItemName  string  `json:"item_name"`
	ItemPrice float64 `json:"item_price"`
	UserID    int64   `json:"user_id"`
	User      User
}

// TableName -> Overriding default table name
func (WishList) TableName() string {
	return "wishlists"
}
