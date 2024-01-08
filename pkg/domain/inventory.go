package domain

// type Inventory struct {
// 	ID          uint     `json:"id" gorm:"unique;not null"`
// 	CategoryID  int      `json:"category_id"`
// 	Category    Category `json:"-" gorm:"foriegnkey:CategoryID;constraint:OnDelete:CASCADE"`
// 	ProductName string   `json:"product_name"`
// 	Size        string   `json:"size" gorm:"size:6;default:'M';check:size IN ('S','M','L','XL','XXL','XXXL')"`
// 	Stock       string   `json:"stock"`
// 	Price       string   `json:"price"`
// }

type Category struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Category string `json:"category"`
}
