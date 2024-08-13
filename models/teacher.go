package models

type Teacher struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Subject   string `json:"subject"`
}

func NewTeacher(firstName string, lastName string, email string, subject string) *Teacher {
	return &Teacher{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Subject:   subject,
	}
}
