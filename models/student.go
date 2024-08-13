package models

type Student struct {
	StudentId int    `json:studentId`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func NewStudent(first string, last string, email string) *Student {
	return &Student{
		FirstName: first,
		LastName:  last,
		Email:     email,
	}
}
