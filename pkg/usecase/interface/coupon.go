package interfaces

import (
	"project/pkg/utils/domain"
	"project/pkg/utils/models"
)

type CouponUseCase interface {
	CreateNewCoupon(coupon models.Coupons) error
	GetAllCoupons() ([]domain.Coupons, error)
	MakeCouponInvalid(couponID int) error
	MakeCouponValid(couponID int) error
}
