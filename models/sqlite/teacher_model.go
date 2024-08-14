package sqlite

import (
	"database/sql"

	"github.com/Yeffian/school_management_api/models"
)

type TeacherModel struct {
	DB *sql.DB
}

func (m *TeacherModel) All() ([]models.Teacher, error) {
	stmt := `SELECT * FROM teachers`
	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	teachers := []models.Teacher{}

	for rows.Next() {
		teacher := models.Teacher{}
		err := rows.Scan(&teacher.TeacherId, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Subject)
		if err != nil {
			return nil, err
		}

		teachers = append(teachers, teacher)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return teachers, nil
}

func (m *TeacherModel) FromFirstName(name string) (*models.Teacher, error) {
	stmt := `SELECT * FROM teachers WHERE firstName = "` + name + `";`
	row := m.DB.QueryRow(stmt)

	t := models.Teacher{}
	row.Scan(&t.TeacherId, &t.FirstName, &t.LastName, &t.Email, &t.Subject)

	err := row.Err()
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (m *TeacherModel) FromLastName(name string) (*models.Teacher, error) {
	stmt := `SELECT * FROM teachers WHERE lastName = "` + name + `";`
	row := m.DB.QueryRow(stmt)

	t := models.Teacher{}
	row.Scan(&t.TeacherId, &t.FirstName, &t.LastName, &t.Email, &t.Subject)

	err := row.Err()
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (m *TeacherModel) FromSubject(subject string) ([]models.Teacher, error) {
	stmt := `SELECT * FROM teachers WHERE subject = "` + subject + `"`
	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	teachers := []models.Teacher{}

	for rows.Next() {
		teacher := models.Teacher{}
		err := rows.Scan(&teacher.TeacherId, &teacher.FirstName, &teacher.LastName, &teacher.Email, &teacher.Subject)
		if err != nil {
			return nil, err
		}

		teachers = append(teachers, teacher)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return teachers, nil
}
