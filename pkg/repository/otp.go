package repository

import (
	interfaces "project/pkg/repository/interface"

	"gorm.io/gorm"
)

type otpRepository struct {
	DB *gorm.DB
}

func NewOtpRepository(DB *gorm.DB) interfaces.OtpRepository {
	return &otpRepository{
		DB: DB,
	}
}
