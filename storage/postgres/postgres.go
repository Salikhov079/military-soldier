package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Salikhov079/military/config"
	u "github.com/Salikhov079/military/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db    *sql.DB
	Soldiers u.Soldier
	Groups u.Group
	Departments u.Department
	Commanders u.Commander
}

func NewPostgresStorage() (u.InitRoot, error) {
	config := config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword,
		config.PostgresHost, config.PostgresPort,
		config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Storage{Db: db}, err
}

func (s *Storage) Soldier() u.Soldier {
	if s.Soldiers == nil {
		s.Soldiers = &SoldierStorage{s.Db}
	}

	return s.Soldiers
}

func (s *Storage) Group() u.Group {
	if s.Groups == nil {
		s.Groups = &GroupStorage{s.Db}
	}

	return s.Groups
}

func (s *Storage) Department() u.Department {
	if s.Departments == nil {
		s.Departments = &DepartmentStorage{s.Db}
	}

	return s.Departments
}

func (s *Storage) Commander() u.Commander {
	if s.Commanders == nil {
		s.Commanders = &CommanderStorage{s.Db}
	}

	return s.Commanders
}