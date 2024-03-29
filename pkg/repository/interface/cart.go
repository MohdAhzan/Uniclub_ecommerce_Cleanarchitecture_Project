package interfaces

type CartRepository interface {
	CreateNewCart(UserID int) (int, error)
	GetCartID(userID int) (int, error)
	AddtoCartItems(cartId, pid int) error
	CheckIfItemIsAlreadyAdded(cartID, pid int) (bool, error)
	UpdateCartQuantity(cartID, pid, quantity int) error
	GetProductIDs(cardTD int) ([]int, error)
	GetProductNames(pID int) (string, error)
	FindCartQuantity(pid, cartID int) (int, error)
	RemoveCartItems(pid, cartID int) error
}
