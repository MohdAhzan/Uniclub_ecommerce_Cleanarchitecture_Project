package models

type AdminLogin struct {
  Email    string `json:"email,omitempty" validate:"required"`
  Password string `json:"password" validate:"min=8,max=20"`
}

type AdminDetailsResponse struct {
  ID    int    `json:"id"`
  Name  string `json:"name"`
  Email string `json:"email"`
}

type GetPaymentMethods struct {
  ID          int    `json:"payment_id"`
  PaymentName string `json:"payment_name"`
}

type SalesReport struct {
  TotalSalesAmount  float64
  TotalOrders       int
  CompletedOrders   int
  PendingOrders     int
  ReturnedOrders    int
  CancelledOrders   int
  MostSelledProduct string
}

type OrderDetailsAdmin struct {
  TotalAmount float64 `gorm:"column:total_amount"`
  ProductName string  `gorm:"column:product_name"`
}

type AdminPasswordChange struct{

  CurrentPassword  string  `json:"current_password"`
  NewPassword      string  `json:"new_password"`
  ConfirmPassword  string  `json:"confirm_password"`

}


