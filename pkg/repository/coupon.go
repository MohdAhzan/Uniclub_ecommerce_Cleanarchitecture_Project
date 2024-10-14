package repository

import (
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/domain"
	"github.com/MohdAhzan/Uniclub_ecommerce_Cleanarchitecture_Project/pkg/utils/models"
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

	err := c.DB.Raw("select count(*) from coupons where coupon_code = ?", couponCode).Scan(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}

	return false, nil
}

func (c *couponRepository) GetAllCoupons() ([]domain.Coupons, error) {

	var couponData []domain.Coupons

	err := c.DB.Raw("select * from coupons").Scan(&couponData).Error
	if err != nil {
		return []domain.Coupons{}, err
	}
	return couponData, nil
}

func (c *couponRepository) MakeCouponInvalid(couponID int) error {

	err := c.DB.Exec("UPDATE coupons SET is_active = FALSE where id = ?", couponID).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *couponRepository) MakeCouponValid(couponID int) error {

	err := c.DB.Exec("UPDATE coupons SET is_active = TRUE where id = ?", couponID).Error
	if err != nil {
		return err
	}
	return nil
}
func (c *couponRepository) CheckCouponStatus(couponID int) (bool, error) {

	var is_active bool

	err := c.DB.Raw("select is_active from coupons where id = ?", couponID).Scan(&is_active).Error
	if err != nil {
		return false, err
	}
	return is_active, nil
}

func (c *couponRepository) FindCouponDetails(couponID int) (domain.Coupons, error) {

	var couponData domain.Coupons

	err := c.DB.Raw("select * from coupons where id = ?", couponID).Scan(&couponData).Error
	if err != nil {
		return domain.Coupons{}, err
	}
	return couponData, nil
}

func (c *couponRepository) CheckIfUserUsedCoupon(userID, couponID int) (bool, error) {

	var count int

	err := c.DB.Raw("select count(*) from orders where user_id = ? and coupon_used_by_id = ? ", userID, couponID).Scan(&count).Error
	if err != nil {
		return true, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
