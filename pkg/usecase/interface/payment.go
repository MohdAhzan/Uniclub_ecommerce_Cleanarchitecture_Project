package interfaces

import "project/pkg/utils/models"

type PaymentUseCase interface {
	MakePaymentFromRazorPay(userID, orderID int) (models.OrderPaymentDetails, error)
	VerifyPaymentFromRazorPay(razorPaymentID, razorOrderID, razorID string) error
}
