package repository

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type paymentRepository struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) *paymentRepository {
	return &paymentRepository{
		DB: DB,
	}
}

func (p *paymentRepository) UpdatePaymentDetails(orderID int) error {
	status := "PAID"

	if err := p.DB.Exec(`UPDATE orders SET payment_status = $1 WHERE id = $2`, status, orderID).Error; err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (p *paymentRepository) InsertPaymentDetails(orderID int, razorID, PaymentID string) error {

	err := p.DB.Exec(`INSERT INTO razor_pays (order_id,razor_id,payment_id) VALUES ($1,$2,$3)`, orderID, razorID, PaymentID).Error
	if err != nil {
		return err
	}
	fmt.Println("inserted Razorpay details ?????????????????????????///")

	return nil
}
