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
	id := uuid.NewString()
	query := `
		INSERT INTO groups (id, name, department_id, commanders_id)
		VALUES ($1, $2, $3, $4)
	`
	_, err := p.db.Exec(query, id, group.Name, group.DepartmentId, group.CommandersId)
	return nil, err
}

func (p *GroupStorage) GetById(id *pb.ById) (*pb.Group, error) {
	query := `
		SELECT id, name, department_id, commanders_id FROM groups 
		WHERE id = $1 and deleted_at = 0
	`
	row := p.db.QueryRow(query, id.Id)

	var group pb.Group
	var department pb.Department
	var commander pb.Commander

	err := row.Scan(
		&group.Id,
		&group.Name,
		&department.Id,
		&commander.Id)
	if err != nil {
		return nil, err
	}
	group.Department = &department
	group.Commanders = &commander

	return &group, nil
}

func (p *GroupStorage) GetAll(filter *pb.GroupReq) (*pb.AllGroups, error) {
	groups := &pb.AllGroups{}
	var arr []interface{}
	count := 1

	query := `SELECT id, name, department_id, commanders_id FROM groups WHERE deleted_at = 0`

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
		var commander pb.Commander

		err := rows.Scan(
			&group.Id,
			&group.Name,
			&department.Id,
			&commander.Id)
		if err != nil {
			return nil, err
		}
		group.Department = &department
		group.Commanders = &commander

		groups.Groups = append(groups.Groups, &group)
	}

	return groups, nil
}

func (p *GroupStorage) Update(group *pb.Group) (*pb.Void, error) {
	query := `
		UPDATE groups
		SET name = $2, department_id = $3, commanders_id = $4, updated_at = now()
		WHERE id = $1 and deleted_at = 0
	`
	_, err := p.db.Exec(query, group.Id, group.Name, group.Department.Id, group.Commanders.Id)
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
