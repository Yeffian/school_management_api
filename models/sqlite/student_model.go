package sqlite

import (
	"database/sql"

	"github.com/Yeffian/school_management_api/models"
)

type StudentModel struct {
	DB *sql.DB
}

func (m *StudentModel) All() ([]models.Student, error) {
	stmt := `SELECT * FROM students`
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

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (m *StudentModel) New(student models.Student) error {
	stmt := `INSERT INTO students VALUES($1, $2, $3, $4, $5)`
	_, err := m.DB.Exec(stmt, student.StudentId, student.FirstName, student.LastName, student.Email, student.ClassCode)

	if err != nil {
		return err
	}

	return nil
}

func (m *StudentModel) ByFirstName(name string) (*models.Student, error) {
	stmt := `SELECT * FROM students WHERE firstName = "` + name + `";`
	row := m.DB.QueryRow(stmt)

	s := models.Student{}
	row.Scan(&s.StudentId, &s.FirstName, &s.LastName, &s.Email, &s.ClassCode)

	err := row.Err()
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (m *StudentModel) ByLastName(name string) (*models.Student, error) {
	stmt := `SELECT * FROM students WHERE lastName = "` + name + `";`
	row := m.DB.QueryRow(stmt)

	s := models.Student{}
	row.Scan(&s.StudentId, &s.FirstName, &s.LastName, &s.Email, &s.ClassCode)

	err := row.Err()
	if err != nil {
		return nil, err
	}

	return &s, nil
}
