package sqlite

import (
	"database/sql" //provides interfaces to work with databases

	_ "github.com/mattn/go-sqlite3"
	"github.com/tejas-2232/students_api/internal/config"
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

	return 0, nil // its returing like ( int64, error)

}
