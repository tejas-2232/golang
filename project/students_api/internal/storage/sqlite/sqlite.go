package sqlite

import (
	"database/sql" //provides interfaces to work with databases
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tejas-2232/students_api/internal/config"
	"github.com/tejas-2232/students_api/internal/types"
)

type Sqlite struct {
	Db *sql.DB // Db is of type:  *sql.DB
}

// go does not have constructors
func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	//create table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	if err != nil {
		return nil, err
	}
	return &Sqlite{
		Db: db,
	}, nil // nil is second return type because Sqlite needs 2 return types
}

// func (s *Pointer_of_struct-Sqlite )
func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students(name, email, age) VALUES(?, ?, ?)") // ? used to prevent sql injection
	// result, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	//binding values using Exec (?,?,?)
	result, err := stmt.Exec(name, email, age)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return lastId, nil // its returing like ( int64, error)

}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM students where id = ? LIMIT 1")

	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student

	// QueryRow is used to query a single row to get the student details
	// scan is used to scan the row into the struct data
	// order matches the order of the columns in the database
	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)

	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)

	}

	return student, nil
}
