package interfaces

type PaymentRepository interface {
	UpdatePaymentDetails(orderID int) error
	InsertPaymentDetails(orderID int, razorID, PaymentID string) error
}
