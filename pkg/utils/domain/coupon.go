package domain

import "time"

type Coupons struct {
	ID           uint      `json:"id" gorm:"primarykey;not null"`
	CouponCode   string    `json:"coupon_code" gorm:"unique;not null"`
	DiscountRate int       `json:"discount_rate" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
	ValidTill    time.Time `json:"valid_till"`
	IsActive     bool      `json:"valid" gorm:"default:true"`
}
