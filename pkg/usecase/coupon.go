package usecase

import (
	"errors"
	interfaces "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/repository/interface"
	domain "github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
	"time"
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

	couponValidDate := coupon.ValidTill.UTC().Truncate(24 * time.Hour)
	currentDate := time.Now().UTC().Truncate(24 * time.Hour)

	if couponValidDate.Before(currentDate) {
		return errors.New("validity should be above present Day")
	}

	err = c.couponRepo.CreateNewCoupon(coupon)
	if err != nil {
		return err
	}

	return nil
}

func (c *couponUseCase) GetAllCoupons() ([]domain.Coupons, error) {

	couponData, err := c.couponRepo.GetAllCoupons()
	if err != nil {
		return []domain.Coupons{}, err
	}

	return couponData, nil

}

func (c *couponUseCase) MakeCouponInvalid(couponID int) error {

	IsActive, err := c.couponRepo.CheckCouponStatus(couponID)
	if err != nil {
		return err
	}
	if !IsActive {
		return errors.New("coupon is already Invalid")
	}

	err = c.couponRepo.MakeCouponInvalid(couponID)
	if err != nil {
		return err
	}
	return nil

}

func (c *couponUseCase) MakeCouponValid(couponID int) error {

	IsActive, err := c.couponRepo.CheckCouponStatus(couponID)
	if err != nil {
		return err
	}
	if IsActive {
		return errors.New("coupon is already Valid")
	}

	err = c.couponRepo.MakeCouponValid(couponID)
	if err != nil {
		return err
	}
	return nil

}
