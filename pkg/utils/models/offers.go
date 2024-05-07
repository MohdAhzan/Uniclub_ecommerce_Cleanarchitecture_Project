package models

import "time"

type AddCategoryOffer struct {
	CategoryID   int       `json:"category_id"`
	OfferName    string    `json:"offer_name"`
	DiscountRate float64   `json:"discount_rate"`
	ValidTill    time.Time `json:"valid_till"`
}

type AddInventoryOffer struct {
	InventoryID  int       `json:"product_id"`
	OfferName    string    `json:"offer_name"`
	DiscountRate float64   `json:"discount_rate"`
	ValidTill    time.Time `json:"valid_till"`
}

type GetInventoryOffers struct {
	ID           uint      `json:"id" gorm:"primarykey;not null"`
	InventoryID  uint      `json:"product_id"`
	ProductName  string    `json:"product_name"`
	OfferName    string    `json:"offer_name"`
	DiscountRate float64   `json:"discount_rate"`
	CreatedAt    time.Time `json:"created_at"`
	ValidTill    time.Time `json:"valid_till"`
	IsActive     bool      `json:"is_active" gorm:"default:false"`
}
