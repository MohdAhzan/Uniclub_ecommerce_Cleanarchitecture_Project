package repository

import "gorm.io/gorm"

type couponRepository struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) *couponRepository {
	return &couponRepository{
		DB: db,
	}
}

func (c *couponRepository) CreateNewCoupon() error {
	return nil
}
