package interfaces

import "project/pkg/utils/models"

type PaymentUseCase interface {
	MakePaymentFromRazorPay(userID, orderID int) (models.OrderPaymentDetails, error)
	VerifyPaymentFromRazorPay(razorPaymentID, razorOrderID, razorID string) error
	PaymentFromWallet(orderID, userID int) (models.OrderPaymentDetails, error)
}
