package models

type Class struct {
	ClassCode string    `json:"classCode"`
	Students  []Student `json:"students"`
	Teachers  []Teacher `json:"teachers"`
}

func NewClass(code string) *Class {
	return &Class{
		ClassCode: code,
	}
}

func (class *Class) AddStudent(student Student) {
	class.Students = append(class.Students, student)
}

func (class *Class) AddTeacher(teacher Teacher) {
	class.Teachers = append(class.Teachers, teacher)
}
