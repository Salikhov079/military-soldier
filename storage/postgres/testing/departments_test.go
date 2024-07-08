package postgres

import (
	"log"
	"testing"

	pb "github.com/Salikhov079/military/genprotos/soldiers"
	"github.com/Salikhov079/military/storage/postgres"

	"github.com/stretchr/testify/assert"
)

func TestCreateDepartment(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	department := &pb.Department{
		Name: "HR Department",
		Commander: &pb.Commander{Id: "92b6362b-b0d1-44b1-9a3e-8b34da4f5b88"},
	}
	result, err := stg.Department().Create(department)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestGetByIdDepartment(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	var id pb.ById
	id.Id = "68ed89b8-63c3-443d-8a2c-1264a563c905"

	department, err := stg.Department().GetById(&id)

	assert.NoError(t, err)
	assert.NotNil(t, department)
}

func TestGetAllDepartments(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}
	departments, err := stg.Department().GetAll(&pb.Department{})
	assert.NoError(t, err)
	assert.NotNil(t, departments)
}

func TestUpdateDepartment(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	department := &pb.Department{
		Id:   "68ed89b8-63c3-443d-8a2c-1264a563c905",
		Name: "Updated HR Department",
		Commander: &pb.Commander{Id: "46705d87-c829-4d2b-9293-2b4b9866e9d0"},
	}
	result, err := stg.Department().Update(department)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestDeleteDepartment(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	var id pb.ById
	id.Id = "80e07ab6-708c-4534-a3b5-9d3e643e78cd"

	result, err := stg.Department().Delete(&id)

	assert.NoError(t, err)
	assert.Nil(t, result)
}
