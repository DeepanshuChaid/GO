package sqlite

import (
	"database/sql"
  "fmt"
	"github.com/DeepanshuChaid/GO/internal/config"
	_ "github.com/mattn/go-sqlite3"
  "github.com/DeepanshuChaid/GO/internal/types"
)



type Sqlite struct {
  Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
  db, err := sql.Open("sqlite3", cfg.StoragePath)
  if err != nil {
    return nil, err
  }

  _, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      name TEXT NOT NULL,
      email TEXT NOT NULL UNIQUE,
      age INTEGER NOT NULL
  )`)
  if err != nil {
    return nil, err
  }
  

  return &Sqlite{
    Db: db,
  }, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
  statement, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
  if err != nil {
    return 0, err
  }
  defer statement.Close()


  result, err := statement.Exec(name, email, age)
  if err != nil {
    return 0, err
  }

  id, err := result.LastInsertId()
  if err != nil {
    return 0, err
  }

  return id, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
  statement, err := s.Db.Prepare(`SELECT * FROM students WHERE id = ? LIMIT 1`)
  if err != nil {
    return types.Student{}, err
  }

  defer statement.Close()

  var student types.Student

  err = statement.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
  if err != nil {
    if err == sql.ErrNoRows {
      return types.Student{}, fmt.Errorf("Student not found")
    }
    
    return types.Student{}, fmt.Errorf("QueryRow failed: %v", err)
  }
  
  return student, nil
}