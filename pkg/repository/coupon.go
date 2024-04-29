package repository

import (
	"project/pkg/utils/models"
	"time"

	"gorm.io/gorm"
)

type couponRepository struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) *couponRepository {
	return &couponRepository{
		DB: db,
	}
}

func (c *couponRepository) CreateNewCoupon(coupon models.Coupons) error {

	err := c.DB.Exec(`INSERT INTO coupons (coupon_code,discount_rate,created_at,valid_till) VALUES (?,?,?,?)`, coupon.CouponCode, coupon.DiscountRate, time.Now(), coupon.ValidTill).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *couponRepository) CheckIfCouponExist(couponCode string) (bool, error) {

	var count int

	err:=c.DB.Raw("select count(*) from coupons where coupon_code = ?",couponCode).Scan(&count).Error
	if err!= nil{
		return false,err
	}
	if count > 0 {
		return true,nil
	}

	return false, nil
}
