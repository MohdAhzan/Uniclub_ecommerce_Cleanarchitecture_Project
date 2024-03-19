package models

type Order struct {
	UserID          int `json:"user_id"`
	AddressID       int `json:"address_id"`
}
