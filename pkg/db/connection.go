package db

import (
	"fmt"
	"project/pkg/config"
	"project/pkg/utils/domain"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	fmt.Println("error creating tables")

	if err := db.AutoMigrate(&domain.Inventories{}); err != nil {

		return db, err
	}
	if err := db.AutoMigrate(&domain.Category{}); err != nil {
		return db, err
	}

	if err := db.AutoMigrate(&domain.Users{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(&domain.Admin{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(&domain.Address{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(&domain.Cart{}); err != nil {
		return db, err
	}
	fmt.Println("cart is created ")
	if err := db.AutoMigrate(&domain.CartItems{}); err != nil {
		return db, err
	}
	if err := db.AutoMigrate(&domain.Order{}); err != nil {
		return db, err
	}

	if err := db.AutoMigrate(&domain.OrderItems{}); err != nil {
		return db, err
	}

	fmt.Println("cartITEms is created ")
	CheckAndCreateAdmin(db)

	return db, dbErr
}

func CheckAndCreateAdmin(db *gorm.DB) {
	var count int64
	db.Model(&domain.Admin{}).Count(&count)
	if count == 0 {
		password := "uniclubadmin"
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return
		}
		admin := domain.Admin{
			ID:       1,
			Name:     "uniclub",
			Email:    "uniclub@gmail.com",
			Password: string(hashedPassword),
		}
		db.Create(&admin)
	}
}
