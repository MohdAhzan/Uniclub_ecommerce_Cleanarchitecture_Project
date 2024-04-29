package models

import "time"

type AddCategoryOffer struct {
	CategoryID   int       `json:"category_id"`
	DiscountRate float64   `json:"discount_rate"`
	ValidTill    time.Time `json:"valid_till"`
}
