package storage

import "github.com/rohan44942/student-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error) // need to give the signature
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudentById(id int64, name string, email string, age int) (types.Student, error)
}
