package models

type Student struct {
	StudentId int    `json:studentId`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	ClassCode string `json:"classCode"`
}

func NewStudent(first string, last string, email string, classCode string) *Student {
	return &Student{
		FirstName: first,
		LastName:  last,
		Email:     email,
		ClassCode: classCode,
	}
}
