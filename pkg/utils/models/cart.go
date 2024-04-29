package models

type AddtoCart struct {
	ProductID int `json:"pid"`
	Quantity  int `json:"quantity"`
}

type CartResponse struct {
	CartID   uint
	CartData []GetCart
}

type GetCart struct {
	ProductID       int     `json:"product_id"`
	ProductName     string  `json:"product_name"`
	Image           string  `json:"image"`
	Category_id     int     `json:"category_id"`
	Quantity        int     `json:"quantity"`
	StockAvailable  int     `json:"stock_available"`
	TotalPrice      float64 `json:"total_price"`
	DiscountedPrice float64 `json:"discounted_price"`
}
