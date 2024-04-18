package models

type OTPData struct {
	PhoneNumber string `json:"phone,omitempty" validate:"required"`
}

type VerifyData struct {
	PhoneNumber string `json:"phone,omitempty" validate:"required"`
	Code        string `json:"code,omitempty" validate:"required"`
}

type EmailTask struct {
	To      string
	Subject string
	Msg     string
}
