package interfaces

import (
	"project/pkg/utils/models"
)

type CouponUseCase interface {
	CreateNewCoupon(coupon models.Coupons) error
}
