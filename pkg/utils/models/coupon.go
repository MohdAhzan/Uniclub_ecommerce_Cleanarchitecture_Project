package models

import "time"

type Coupons struct {
	CouponCode   string    `json:"coupon_code" gorm:"unique;not null"`
	DiscountRate int       `json:"discount_rate" gorm:"not null"`
	ValidTill    time.Time `json:"valid_till"`
	IsActive     bool      `json:"valid" gorm:"default:true"`
}
