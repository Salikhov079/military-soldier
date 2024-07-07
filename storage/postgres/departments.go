package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Salikhov079/military/genprotos/soldiers"

	"github.com/google/uuid"
)

type DepartmentStorage struct {
	db *sql.DB
}

func NewDepartmentStorage(db *sql.DB) *DepartmentStorage {
	return &DepartmentStorage{db: db}
}

func (p *DepartmentStorage) Create(department *pb.Department) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO departments (id, name)
		VALUES ($1, $2)
	`
	_, err := p.db.Exec(query, id, department.Name)
	return nil, err
}

func (p *DepartmentStorage) GetById(id *pb.ById) (*pb.Department, error) {
	query := `
		SELECT id, name FROM departments 
		WHERE id = $1 and deleted_at = 0
	`
	row := p.db.QueryRow(query, id.Id)

	var department pb.Department

	err := row.Scan(
		&department.Id,
		&department.Name)
	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (p *DepartmentStorage) GetAll(filter *pb.Department) (*pb.AllDepartments, error) {
	departments := &pb.AllDepartments{}
	var arr []interface{}
	count := 1

	query := `SELECT id, name FROM departments WHERE deleted_at = 0`

	if len(filter.Name) > 0 {
		query += fmt.Sprintf(" AND name=$%d", count)
		count++
		arr = append(arr, filter.Name)
	}

	rows, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var department pb.Department

		err = rows.Scan(
			&department.Id,
			&department.Name)
		if err != nil {
			return nil, err
		}

		departments.Departments = append(departments.Departments, &department)
	}

	return departments, nil
}

func (p *DepartmentStorage) Update(department *pb.Department) (*pb.Void, error) {
	query := `
		UPDATE departments
		SET name = $2, updated_at = now()
		WHERE id = $1 and deleted_at = 0
	`
	_, err := p.db.Exec(query, department.Id, department.Name)
	return nil, err
}

func (p *DepartmentStorage) Delete(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE departments SET deleted_at = $2
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, id.Id, time.Now().Unix())
	return nil, err
}