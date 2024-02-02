package models

type Inventories struct {
	Product_ID  uint    `json:"product_id"`
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size" `
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type AddInventory struct {
	Product_ID  uint    `json:"product_id"`
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
