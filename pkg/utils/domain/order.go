package domain

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID              uint           `gorm:"primarykey"`
	CreatedAt       time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	UserID          uint           `json:"user_id" gorm:"not null"`
	Users           Users          `json:"-" gorm:"foreignkey:UserID"`
	AddressID       uint           `json:"address_id" gorm:"not null"`
	Address         Address        `json:"-" gorm:"foreignkey:AddressID"`
	PaymentMethodID uint           `json:"paymentmethod_id"`
	PaymentMethod   PaymentMethod  `json:"-" gorm:"foreignkey:PaymentMethodID"`
	CouponUsedByID  int            `json:"coupon_used_by_id" gorm:"default:null"`
	Coupons         Coupons        `json:"-" gorm:"foreignkey:CouponUsedByID"`
	FinalPrice      float64        `json:"price"`
	OrderStatus     string         `json:"order_status" gorm:"order_status:5;default:'PENDING';check:order_status IN ('PENDING', 'SHIPPED','DELIVERED','CANCELED','RETURN_REQUESTED','RETURNED')"`
	PaymentStatus   string         `json:"payment_status" gorm:"payment_status:2;default:'NOT PAID';check:payment_status IN ('PAID', 'NOT PAID')"`
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
	AddressID     uint    `json:"address_id"`
	Address       Address `json:"address"`
	OrderStatus   string  `json:"order_status"`
	PaymentMethod string  `json:"payment_method" gorm:"payment_method"`
	PaymentStatus string  `json:"payment_status"`
	Total         float64 `json:"total"`
}
type OrderDetailsSeparate struct {
	ID       int    `json:"id" gorm:"id"`
	Username string `json:"name"`
	// AddressID               uint                      `json:"address_id"`
	Address                 Address                   `json:"address"`
	OrderStatus             string                    `json:"order_status"`
	PaymentMethod           string                    `json:"payment_method" gorm:"payment_method"`
	PaymentStatus           string                    `json:"payment_status"`
	EachProductOrderDetails []EachProductOrderDetails `json:"product_order_details"`
	Total                   float64                   `json:"total"`
}

// type EachProductOrderDetails struct {
// 	ProductID     uint    `json:"product_id"`
// 	Quantity      int    `json:"total_quantity"`
// 	ProductPrice  float64 `json:"total_price"`
// 	DiscountPrice float64 `json:"discount_price"`
// 	ProductOffer  string  `json:"offer_name"`
// }

type EachProductOrderDetails struct {
	ProductID     uint    `json:"product_id"`
	Quantity      uint    `json:"total_quantity"`
	ProductPrice  float64 `json:"total_price"`
	DiscountPrice float64 `json:"discount_price"`
	ProductOffer  string  `json:"offer_name"`
}
type OrderDetailsWithImages struct {
	OrderDetails  Order
	ProductID     []int
	Images        []string
	PaymentMethod string
}

type AdminOrdersResponse struct {
	PENDING          []AdminOrderDetails
	SHIPPED          []AdminOrderDetails
	DELIVERED        []AdminOrderDetails
	CANCELED         []AdminOrderDetails
	RETURN_REQUESTED []AdminOrderDetails
	RETURNED         []AdminOrderDetails
}

type AdminOrderDetails struct {
	ID            int       `json:"order_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	UserID        int       `json:"user_id"`
	Name          string    `json:"user_name"`
	Address       string    `json:"address"`
	Landmark      string    `json:"land_mark"`
	City          string    `json:"city"`
	Pincode       string    `json:"pincode"`
	State         string    `json:"state"`
	Phone         string    `json:"phone"`
	Default       bool      `json:"default"`
	PaymentMethod string    `json:"payment_method"`
	Total         float64   `json:"total"`
	OrderStatus   string    `json:"order_status"`
	PaymentStatus string    `json:"payment_status"`
}

type PaymentMethod struct {
	ID          uint   `gorm:"primarykey"`
	PaymentName string `json:"payment_name"`
	IsDeleted   bool   `json:"is_deleted" gorm:"default:false"`
}
