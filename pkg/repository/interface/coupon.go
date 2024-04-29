package interfaces

import "project/pkg/utils/models"

type CouponRepository interface {
	CreateNewCoupon(coupon models.Coupons) error
	CheckIfCouponExist(couponCode string) (bool, error)
}
