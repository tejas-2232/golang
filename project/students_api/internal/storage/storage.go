package storage

import "github.com/tejas-2232/students_api/internal/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error) // returning slice of Student struct
	UpdateStudentById(id int64, name string, email string, age int) error
	DeleteStudentById(id int64) error
	// implement this methid in sqlite.go

}
