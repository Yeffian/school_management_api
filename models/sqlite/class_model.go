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
