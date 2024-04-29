package usecase

import (
	"errors"
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

	exist, err := c.couponRepo.CheckIfCouponExist(coupon.CouponCode)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("coupon already exist in this name")
	}

	err = c.couponRepo.CreateNewCoupon(coupon)
	if err != nil {
		return err
	}

	return nil
}
