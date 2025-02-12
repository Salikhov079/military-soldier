package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	pb "github.com/Salikhov079/military/genprotos/soldiers"

	"github.com/google/uuid"
)

type SoldierStorage struct {
	db *sql.DB
}

func NewSoldierStorage(db *sql.DB) *SoldierStorage {
	return &SoldierStorage{db: db}
}

func (p *SoldierStorage) Create(soldier *pb.SoldierReq) (*pb.Void, error) {
	err := p.checkAge(soldier)
	if err != nil {
		return nil, err
	}
	var count int
	err = p.db.QueryRow("SELECT COUNT(1) FROM soldiers WHERE group_id = $1 GROUP BY group_id", soldier.GroupId).Scan(&count)
	if err != nil && err != sql.ErrNoRows || count == 10 {
		return nil, errors.New("error group is full")
	}
	id := uuid.NewString()
	query := `
		INSERT INTO soldiers (id, name, email, date_of_birth, phone_number, group_id, join_date, end_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err = p.db.Exec(query, id, soldier.Name, soldier.Email, soldier.DateOfBirth, soldier.PhoneNumber, soldier.GroupId, soldier.JoinDate, soldier.EndDate)
	return nil, err
}

func (p *SoldierStorage) GetById(id *pb.ById) (*pb.Soldier, error) {
	query := `
		SELECT id, name, email, date_of_birth, phone_number, group_id, join_date, end_date FROM soldiers 
		WHERE id = $1 and deleted_at = 0
	`
	row := p.db.QueryRow(query, id.Id)

	var soldier pb.Soldier
	var group pb.Group

	err := row.Scan(
		&soldier.Id,
		&soldier.Name,
		&soldier.Email,
		&soldier.DateOfBirth,
		&soldier.PhoneNumber,
		&group.Id,
		&soldier.JoinDate,
		&soldier.EndDate)
	if err != nil {
		return nil, err
	}
	soldier.Group = &group

	return &soldier, nil
}

func (p *SoldierStorage) GetAll(filter *pb.SoldierReq) (*pb.AllSoldiers, error) {
	soldiers := &pb.AllSoldiers{}
	var arr []interface{}
	count := 1

	query := `SELECT id, name, email, date_of_birth, phone_number, group_id, join_date, end_date FROM soldiers WHERE deleted_at = 0`

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

	if len(filter.DateOfBirth) > 0 {
		query += fmt.Sprintf(" AND EXTRACT(YEAR FROM AGE(date_of_birth)) =$%d", count)
		count++
		arr = append(arr, filter.DateOfBirth)
	}

	if len(filter.JoinDate)>0 {
		query+= fmt.Sprintf(` AND EXTRACT(MONTH FROM join_date) = EXTRACT(MONTH FROM $%d::date)
			AND EXTRACT(YEAR FROM join_date) = EXTRACT(YEAR FROM $%d::date);`, count, count)
		count++
		arr=append(arr, filter.JoinDate)
	}

	if len(filter.EndDate)>0 {
		query+= fmt.Sprintf(` AND EXTRACT(MONTH FROM end_date) = EXTRACT(MONTH FROM $%d::date)
			AND EXTRACT(YEAR FROM end_date) = EXTRACT(YEAR FROM $%d::date);`, count, count)
		count++
		arr=append(arr, filter.EndDate)
	}
	

	rows, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var soldier pb.Soldier
		var group pb.Group

		err = rows.Scan(
			&soldier.Id,
			&soldier.Name,
			&soldier.Email,
			&soldier.DateOfBirth,
			&soldier.PhoneNumber,
			&group.Id,
			&soldier.JoinDate,
			&soldier.EndDate)
		if err != nil {
			return nil, err
		}
		soldier.Group = &group

		soldiers.Soldiers = append(soldiers.Soldiers, &soldier)
	}

	return soldiers, nil
}

func (p *SoldierStorage) Update(soldier *pb.Soldier) (*pb.Void, error) {
	query := `
		UPDATE soldiers
		SET name = $2, email = $3, date_of_birth = $4, phone_number = $5, group_id = $6, join_date = $7, end_date = $8, updated_at = now()
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, soldier.Id, soldier.Name, soldier.Email, soldier.DateOfBirth, soldier.PhoneNumber, soldier.Group.Id, soldier.JoinDate, soldier.EndDate)
	return nil, err
}

func (p *SoldierStorage) Delete(id *pb.ById) (*pb.Void, error) {
	query := `
		UPDATE soldiers SET deleted_at = $2
		WHERE id = $1 AND deleted_at = 0
	`
	_, err := p.db.Exec(query, id.Id, time.Now().Unix())
	return nil, err
}

func (p *SoldierStorage) UseBullet(use *pb.UseB) (*pb.Void, error) {
	query := `
		SELECT quantity_weapon, quantity_big_weapon FROM use_bullets
		WHERE date = now()::date and soldier_id = $1
	`
	rows, err := p.db.Query(query, use.SoldierId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data1, data2 int

	for rows.Next() {
		var d1, d2 int
		err = rows.Scan(&d1, &d2)
		if err != nil {
			return nil, err
		}
		data1 += d1
		data2 += d2
	}

	data1 += int(use.QuantityWeapon)
	if data1 > 50 {
		return nil, errors.New("error: don't use more then 50 Weapons in 1 day")
	}
	data2 += int(use.QuantityBigWeapon)
	if data2 > 5 {
		return nil, errors.New("error: don't use more then 5 Big Weapons in 1 day")
	}

	id := uuid.NewString()
	query = `
		INSERT INTO use_bullets (id, quantity_weapon, quantity_big_weapon, soldier_id, date)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = p.db.Exec(query, id, use.QuantityWeapon, use.QuantityBigWeapon, use.SoldierId, use.Date)
	return &pb.Void{}, err
}

func (p *SoldierStorage) UseFuel(use *pb.UseF) (*pb.Void, error) {
	query := `
		SELECT diesel, petrol FROM use_fuels
		WHERE date = now()::date and soldier_id = $1
	`
	rows, err := p.db.Query(query, use.SoldierId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data1, data2 int

	for rows.Next() {
		var d1, d2 int
		err = rows.Scan(&d1, &d2)
		if err != nil {
			return nil, err
		}
		data1 += d1
		data2 += d2
	}

	data1 += int(use.Diesel)
	if data1 > 50 {
		return nil, errors.New("error: don't use more then 50l disel in 1 day")
	}
	data2 += int(use.Petrol)
	if data2 > 50 {
		return nil, errors.New("error: don't use more then 50l petrol in 1 day")
	}

	id := uuid.NewString()
	query = `
		INSERT INTO use_fuels (id, diesel, petrol, soldier_id, date)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err = p.db.Exec(query, id, use.Diesel, use.Petrol, use.SoldierId, use.Date)
	return &pb.Void{}, err
}

func (p *SoldierStorage) checkAge(soldier *pb.SoldierReq) error {
	dob, err := time.Parse(time.RFC3339, soldier.DateOfBirth)
	if err != nil {
		return err
	}
	now := time.Now()
	ageYears := now.Year() - dob.Year()
	if now.Before(dob.AddDate(ageYears, 0, 0)) {
		ageYears--
	}
	if ageYears < 18 {
		return errors.New("invalid age: age is less than 18")
	}
	return nil
}

func (p *SoldierStorage) GetAllWeaponStatistik(filter *pb.GetSoldierStatistik) (*pb.GetSoldierStatistikRes, error) {
	soldiers := &pb.GetSoldierStatistikRes{}
	var arr []interface{}

	query := `
        SELECT soldier_id, SUM(quantity_weapon) AS total_quantity_weapon, 
    SUM(quantity_big_weapon) AS total_quantity_big_weapon
        FROM use_bullets
        WHERE date = $1
    `
	arr = append(arr, filter.Date)

	if len(filter.SoldierId) > 0 {
		query += ` AND soldier_id = $2`
		arr = append(arr, filter.SoldierId)
	}

	query += ` GROUP BY soldier_id`

	rows, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var weapons pb.UseB
		err = rows.Scan(&weapons.SoldierId, &weapons.QuantityWeapon, &weapons.QuantityBigWeapon)
		if err != nil {
			return nil, err
		}
		soldiers.UsedWeapons = append(soldiers.UsedWeapons, &weapons)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return soldiers, nil
}

func (p *SoldierStorage) GetAllFuelStatistik(filter *pb.GetSoldierStatistikFuel) (*pb.GetSoldierStatistikFuelRes, error) {
	soldiers := &pb.GetSoldierStatistikFuelRes{}
	var arr []interface{}

	query := `
        SELECT soldier_id, SUM(diesel) AS total_diesel, SUM(petrol) AS total_petrol
        FROM use_fuels
        WHERE date = $1
    `
	arr = append(arr, filter.Date)

	if len(filter.SoldierId) > 0 {
		query += ` AND soldier_id = $2`
		arr = append(arr, filter.SoldierId)
	}

	query += ` GROUP BY soldier_id`

	rows, err := p.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var fuel pb.UseF
		err = rows.Scan(&fuel.SoldierId, &fuel.Diesel, &fuel.Petrol)
		if err != nil {
			return nil, err
		}
		soldiers.UsedFuel = append(soldiers.UsedFuel, &fuel)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return soldiers, nil
}
