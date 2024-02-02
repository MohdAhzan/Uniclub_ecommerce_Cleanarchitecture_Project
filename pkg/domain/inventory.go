package domain

type Inventories struct {
	Product_ID  uint     `json:"product_id" gorm:"primaryKey;not null;autoIncrement"`
	CategoryID  int      `json:"category_id"`
	Category    Category `json:"-" gorm:"foriegnkey:CategoryID;constraint:OnDelete:CASCADE"`
	ProductName string   `json:"product_name"`
	Size        string   `json:"size" gorm:"size:6;default:'M';check:size IN ('S','M','L','XL','XXL','XXXL')"`
	Stock       int      `json:"stock"`
	Price       float64  `json:"price"`
}

type Category struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Category string `json:"category"`
}
