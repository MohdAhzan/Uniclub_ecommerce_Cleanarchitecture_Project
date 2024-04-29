package models

import "time"

type Coupons struct {
	CouponCode   string    `json:"coupon_code"`
	DiscountRate int       `json:"discount_rate"`
	ValidTill    time.Time `json:"valid_till"`
}
