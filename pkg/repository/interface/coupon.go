package interfaces

import (
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
)

type CouponRepository interface {
	CreateNewCoupon(coupon models.Coupons) error
	CheckIfCouponExist(couponCode string) (bool, error)
	GetAllCoupons() ([]domain.Coupons, error)
	MakeCouponInvalid(couponID int) error
	MakeCouponValid(couponID int) error
	CheckCouponStatus(couponID int) (bool, error)
	FindCouponDetails(couponID int) (domain.Coupons, error)
	CheckIfUserUsedCoupon(userID, couponID int) (bool, error)
}
