package models

type Inventories struct {
	Product_ID  uint    `json:"product_id"`
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Image       string  `json:"image"`
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

type EditInventory struct {
	CategoryID  int     `json:"category_id"`
	ProductName string  `json:"product_name"`
	Size        string  `json:"size"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
