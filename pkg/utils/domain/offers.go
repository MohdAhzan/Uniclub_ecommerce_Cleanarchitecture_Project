package domain

import "time"

type CategoryOffers struct {
	ID           uint      `json:"id" gorm:"primarykey;not null"`
	CategoryID   uint      `json:"category_id"`
	Category     Category  `json:"-" gorm:"foreignkey:CategoryID;constraint:OnDelete:CASCADE"`
	DiscountRate float64   `json:"discount_rate"`
	CreatedAt    time.Time `json:"created_at"`
	ValidTill    time.Time `json:"valid_till"`
	IsActive     bool      `json:"is_active" gorm:"default:false"`
}

type InventoryOffers struct {
	ID           uint        `json:"id" gorm:"primarykey;not null"`
	InventoryID  uint        `json:"category_id"`
	Inventories  Inventories `json:"-" gorm:"foreignkey:InventoryID;constraint:OnDelete:CASCADE"`
	DiscountRate float64     `json:"discount_rate"`
	CreatedAt    time.Time   `json:"created_at"`
	ValidTill    time.Time   `json:"valid_till"`
	IsActive     bool        `json:"is_active" gorm:"default:false"`
}
