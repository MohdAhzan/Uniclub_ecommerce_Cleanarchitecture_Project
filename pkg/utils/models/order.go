package models

type Order struct {
	UserID    int `json:"user_id"`
	AddressID int `json:"address_id"`
	PaymentID int `json:"payment_id"`
}

type OrderData struct {
	Payment_method_id int
	Order_status      string
	Final_Price       float64
	PaymentStatus     string
}

type OrderPaymentDetails struct {
	UserID     int     `json:"user_id"`
	Username   string  `json:"username"`
	Razor_id   string  `josn:"razor_id"`
	OrderID    int     `json:"order_id"`
	FinalPrice float64 `json:"final_price"`
}

type EachOrderData struct {
	ProductID     int     `json:"product_id"`
	TotalQuantity int     `json:"total_quantity"`
	TotalPrice    float64 `json:"total_price"`
}
