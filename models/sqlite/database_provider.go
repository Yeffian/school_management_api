package sqlite

import "database/sql"

type DatabaseProvider struct {
	Student StudentModel
	Teacher TeacherModel
	Class   ClassModel
}

func CreateDatabaseProvider(db *sql.DB) *DatabaseProvider {
	return &DatabaseProvider{
		Student: StudentModel{DB: db},
		Teacher: TeacherModel{DB: db},
		Class:   ClassModel{DB: db},
	}
}
