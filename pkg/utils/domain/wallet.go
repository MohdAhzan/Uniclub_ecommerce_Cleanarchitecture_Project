package domain

type Wallet struct {
	WalletID        uint    `json:"wallet_id" gorm:"primarykey;autoIncrement"`
	UserID          uint    `json:"user_id"`
	Users           Users   `json:"-" gorm:"foreignkey : UserID"`
	WalletAmount    float64 `json:"wallet_amount" gorm:"default:0"`
	TranscationType string  `json:"transaction_type" gorm:"transaction_type:3;check:transaction_type IN ('REFERAL','PDT_CANCELLED','PDT_RETURNED')"`
}
