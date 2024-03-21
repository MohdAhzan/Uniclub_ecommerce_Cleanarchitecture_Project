package interfaces

type OrderRepository interface {
	OrderItems(userID, address int, price float64) (int, error)
}
