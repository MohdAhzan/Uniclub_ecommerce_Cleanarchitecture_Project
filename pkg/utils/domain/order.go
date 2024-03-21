package domain

import "gorm.io/gorm"

type PaymentMethod struct {
	ID           uint   `gorm:"primarykey"`
	Payment_Name string `json:"payment_name"`
	IsDeleted    bool   `json:"is_deleted" gorm:"default:false"`
}

type Order struct {
	gorm.Model
	UserID    uint    `json:"user_id" gorm:"not null"`
	Users     Users   `json:"-" gorm:"foreignkey:UserID"`
	AddressID uint    `json:"address_id" gorm:"not null"`
	Address   Address `json:"-" gorm:"foreignkey:AddressID"`
	// PaymentMethodID uint          `json:"paymentmethod_id"`
	// PaymentMethod   PaymentMethod `json:"-" gorm:"foreignkey:PaymentMethodID"`
	Price         float64 `json:"price"`
	OrderStatus   string  `json:"order_status" gorm:"order_status:4;default:'PENDING';check:order_status IN ('PENDING', 'SHIPPED','DELIVERED','CANCELED','RETURNED')"`
	PaymentStatus string  `json:"payment_status" gorm:"payment_status:2;default:'NOT PAID';check:payment_status IN ('PAID', 'NOT PAID')"`
}

type OrderItems struct {
	ID          uint        `json:"id" gorm:"primarykey;autoIncrement"`
	OrderID     uint        `json:"order_id"`
	Order       Order       `json:"-" gorm:"foreignkey:OrderID;constraint:OnDelete:CASCADE"`
	InventoryID uint        `json:"inventory_id"`
	Inventories Inventories `json:"-" gorm:"foreignkey:InventoryID"`
	Quantity    int         `json:"quantity"`
	TotalPrice  float64     `json:"total_price"`
}
