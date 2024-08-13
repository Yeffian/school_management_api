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
		err := rows.Scan(&student.StudentId, &student.FirstName, &student.LastName, &student.Email)
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
