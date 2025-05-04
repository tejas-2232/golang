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

func (s *Sqlite) GetStudents() ([]types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id, name, email, age FROM students")

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query() // Query returns a rows and error

	if err != nil {
		return nil, err
	}
	defer rows.Close() // closing rows

	var students []types.Student // sytoring records in students

	for rows.Next() {
		var student types.Student

		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Age)

		if err != nil {
			return nil, err
		}
		// if everything is right then append the student to the students slice
		// append is a built-in function that adds the student to the end of the students slice.
		// it returns a new slice that contains the original slice and the new element.
		// the new slice is then assigned back to the students variable.
		// this is how we build up the list of students in the students slice.
		students = append(students, student)
	}
	return students, nil
}

func (s *Sqlite) UpdateStudentById(id int64, name string, email string, age int) error {
	stmt, err := s.Db.Prepare("UPDATE students SET name = ?, email = ?, age = ? WHERE id = ?")

	if err != nil {
		return err
	}
	defer stmt.Close()
	//execute the update query
	result, err := stmt.Exec(name, email, age, id)

	if err != nil {
		return err
	}

	//check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no student found with id %s", fmt.Sprint(id))
	}
	return nil
}

// delete student by id
func (s *Sqlite) DeleteStudentById(id int64) error {
	stmt, err := s.Db.Prepare("DELETE FROM students WHERE id = ?")

	if err != nil {
		return err
	}

	defer stmt.Close()

	//execute the delete query
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	//check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no student found with id %s", fmt.Sprint(id))
	}
	return nil
}

// close the database connection
func (s *Sqlite) Close() error {
	return s.Db.Close()
}
