package models

type Class struct {
	ClassCode string `json:"classCode"`
	Subject   string `json:"subject"`
}

func NewClass(code string, subject string) *Class {
	return &Class{
		ClassCode: code,
		Subject:   subject,
	}
}
