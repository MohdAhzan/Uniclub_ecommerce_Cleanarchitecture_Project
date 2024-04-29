package usecase

import (
	interfaces "project/pkg/repository/interface"
	"project/pkg/utils/models"
)

type couponUseCase struct {
	couponRepo interfaces.CouponRepository
}

func NewCouponUseCase(coup interfaces.CouponRepository) *couponUseCase {
	return &couponUseCase{
		couponRepo: coup,
	}
}

func (c *couponUseCase) CreateNewCoupon(coupon models.Coupons) error {
	return nil
}
