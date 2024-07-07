package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/Salikhov079/military/genprotos/soldiers"

	"github.com/google/uuid"
)

type CommanderStorage struct {
	db *sql.DB
}

func NewCommanderStorage(db *sql.DB) *CommanderStorage {
	return &CommanderStorage{db: db}
}

func (p *CommanderStorage) Create(commander *pb.CommanderReq) (*pb.Void, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO commanders (id, name, email, date_of_birth, phone_number, position)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := p.db.Exec(query, id, commander.Name, commander.Email, commander.DateOfBirth, commander.PhoneNumber, commander.Position)
	return nil, err
}

func (p *CommanderStorage) GetById(id *pb.ById) (*pb.Commander, error) {
	query := `
		SELECT id, name, email, date_of_birth, phone_number, position FROM commanders 
		WHERE id = $1 AND deleted_at = 0
	`
	row := p.db.QueryRow(query, id.Id)

	var commander pb.Commander

	err := row.Scan(
		&commander.Id,
		&commander.Name,
		&commander.Email,
		&commander.DateOfBirth,
		&commander.PhoneNumber,
		&commander.Position)
	if err != nil {
		return nil, err
	}

	return &commander, nil
}

func (p *CommanderStorage) GetAll(filter *pb.CommanderReq) (*pb.AllCommanders, error) {
	commanders := &pb.AllCommanders{}
	var arr []interface{}
	count := 1

	query := `SELECT id, name, email, date_of_birth, phone_number, position FROM commanders WHERE deleted_at = 0`

	if len(filter.Name) > 0 {
		query += fmt.Sprintf(" AND name=$%d", count)
		count++
		arr = append(arr, filter.Name)
	}

	if len(filter.Email) > 0 {
		query += fmt.Sprintf(" AND email=$%d", count)
		count++
		arr = append(arr, filter.Email)
	}

	rows, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var commander pb.Commander

		err = rows.Scan(
			&commander.Id,
			&commander.Name,
			&commander.Email,
			&commander.DateOfBirth,
			&commander.PhoneNumber,
			&commander.Position)
		if err != nil {
			return nil, err
		}

		commanders.Commanders = append(commanders.Commanders, &commander)
	}

	return commanders, nil
}

func (p *CommanderStorage) Update(commander *pb.Commander) (*pb.Void, error) {
	query := `
		UPDATE commanders
		SET name = $2, email = $3, date_of_birth = $4, phone_number = $5, position = $6, updated_at = now()
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, commander.Id, commander.Name, commander.Email, commander.DateOfBirth, commander.PhoneNumber, commander.Position)
	return nil, err
}

func (p *CommanderStorage) Delete(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE commanders SET deleted_at = $2
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, id.Id, time.Now().Unix())
	return nil, err
}
