package domain

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID        uint    `json:"user_id" gorm:"not null"`
	Users         Users   `json:"-" gorm:"foreignkey:UserID"`
	AddressID     uint    `json:"address_id" gorm:"not null"`
	Address       Address `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethod string  `json:"payment_method" gorm:"default:'Cash on Delivery'"`
	Price         float64 `json:"price"`
	OrderStatus   string  `json:"order_status" gorm:"order_status:5;default:'PENDING';check:order_status IN ('PENDING', 'SHIPPED','DELIVERED','CANCELED','RETURN_REQUESTED','RETURNED')"`
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

type OrderDetails struct {
	ID            int     `json:"id" gorm:"id"`
	Username      string  `json:"name"`
	Address       Address `json:"address"`
	OrderStatus   string  `json:"order_status"`
	PaymentMethod string  `json:"payment_method" gorm:"payment_method"`
	PaymentStatus string  `json:"payment_status"`
	Total         float64 `json:"total"`
}

type OrderDetailsWithImages struct {
	OrderDetails  Order
	Images        []string
	PaymentMethod string
}
