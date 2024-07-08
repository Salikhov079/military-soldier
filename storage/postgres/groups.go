package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Salikhov079/military/genprotos/soldiers"

	"github.com/google/uuid"
)

type GroupStorage struct {
	db *sql.DB
}

func NewGroupStorage(db *sql.DB) *GroupStorage {
	return &GroupStorage{db: db}
}

func (p *GroupStorage) Create(group *pb.GroupReq) (*pb.Void, error) {
	err := p.checkPostion(group)
	if err != nil {
		return nil, err
	}
	id := uuid.NewString()
	query := `
		INSERT INTO groups (id, name, department_id)
		VALUES ($1, $2, $3)
	`
	_, err = p.db.Exec(query, id, group.Name, group.DepartmentId)
	return nil, err
}

func (p *GroupStorage) GetById(id *pb.ById) (*pb.Group, error) {
	query := `
		SELECT id, name, department_id FROM groups 
		WHERE id = $1 and deleted_at = 0
	`
	row := p.db.QueryRow(query, id.Id)

	var group pb.Group
	var department pb.Department

	err := row.Scan(
		&group.Id,
		&group.Name,
		&department.Id)
	if err != nil {
		return nil, err
	}
	group.Department = &department

	return &group, nil
}

func (p *GroupStorage) GetAll(filter *pb.GroupReq) (*pb.AllGroups, error) {
	groups := &pb.AllGroups{}
	var arr []interface{}
	count := 1

	query := `SELECT id, name, department_id FROM groups WHERE deleted_at = 0`

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
		var group pb.Group
		var department pb.Department

		err := rows.Scan(
			&group.Id,
			&group.Name,
			&department.Id)
		if err != nil {
			return nil, err
		}
		group.Department = &department

		groups.Groups = append(groups.Groups, &group)
	}

	return groups, nil
}

func (p *GroupStorage) Update(group *pb.Group) (*pb.Void, error) {
	query := `
		UPDATE groups
		SET name = $2, department_id = $3, updated_at = now()
		WHERE id = $1 and deleted_at = 0
	`
	_, err := p.db.Exec(query, group.Id, group.Name, group.Department.Id)
	return nil, err
}

func (p *GroupStorage) Delete(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE groups SET deleted_at = $2
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, id.Id, time.Now().Unix())
	return nil, err
}

func (p *GroupStorage) checkPostion(group *pb.GroupReq) error {
	var count int
	err := p.db.QueryRow("SELECT COUNT(1) FROM groups WHERE department_id = $1 GROUP BY department_id", group.DepartmentId).Scan(&count)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	var com string
	err = p.db.QueryRow("SELECT commanders_id FROM departments WHERE id = $1", group.DepartmentId).Scan(&com)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if count == 10 {
		_, err = p.db.Exec("UPDATE commander SET position = 'yuzboshi' WHERE id = $1", com)
		if err != nil {
			return err
		}
	}
	if count == 100 {
		_, err = p.db.Exec("UPDATE commander SET position = 'mingboshi' WHERE id = $1", com)
		if err != nil {
			return err
		}
	}
	if count == 1000 {
		_, err = p.db.Exec("UPDATE commander SET position = 'amir' WHERE id = $1", com)
		if err != nil {
			return err
		}
	}
	if count == 10000 {
		_, err = p.db.Exec("UPDATE commander SET position = 'qomondon' WHERE id = $1", com)
		if err != nil {
			return err
		}
	}
	return nil
}
