package models

type Teacher struct {
	TeacherId int    `json:teacherId`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Subject   string `json:"subject"`
	ClassCode string `json:"classCode"`
}

func NewTeacher(firstName string, lastName string, email string, subject string, classCode string) *Teacher {
	return &Teacher{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Subject:   subject,
		ClassCode: classCode,
	}
}
