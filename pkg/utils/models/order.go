package models

type Order struct {
	UserID    int `json:"user_id"`
	AddressID int `json:"address_id"`
}

type OrderData struct {
	Payment_method string
	Order_status   string
	Price          float64
	PaymentStatus  string
}

