package sqlite

import (
	"database/sql"

	"github.com/Yeffian/school_management_api/models"
)

type ClassModel struct {
	DB *sql.DB
}

func (m *ClassModel) All() ([]models.Class, error) {
	stmt := `SELECT * FROM classes`
	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	classes := []models.Class{}

	for rows.Next() {
		class := models.Class{}
		err := rows.Scan(&class.ClassCode, &class.Subject)
		if err != nil {
			return nil, err
		}

		classes = append(classes, class)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (m *ClassModel) TeachersByClass(class string) ([]models.Teacher, error) {
	stmt := `SELECT * FROM teachers WHERE classCode = "` + class + `";`
	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	teachers := []models.Teacher{}

	for rows.Next() {
		teacher := models.Teacher{}
		err := rows.Scan(&teacher.TeacherId, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Subject, &teacher.ClassCode)
		if err != nil {
			return nil, err
		}

		teachers = append(teachers, teacher)
	}

	return teachers, nil
}

func (m *ClassModel) StudentsByClass(class string) ([]models.Student, error) {
	stmt := `SELECT * FROM students WHERE classCode = "` + class + `";`
	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	students := []models.Student{}

	for rows.Next() {
		student := models.Student{}
		err := rows.Scan(&student.StudentId, &student.FirstName, &student.LastName, &student.Email, &student.ClassCode)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}
