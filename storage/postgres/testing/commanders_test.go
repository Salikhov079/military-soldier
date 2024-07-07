package postgres

import (
	"log"
	"testing"

	pb "github.com/Salikhov079/military/genprotos/soldiers"
	"github.com/Salikhov079/military/storage/postgres"

	"github.com/stretchr/testify/assert"
)

func TestCreateCommander(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	commander := &pb.CommanderReq{
		Name:        "John Doe",
		Email:       "john.doe@example.com",
		DateOfBirth: "1980-01-01",
		PhoneNumber: "+1234567890",
		Position:    "o'nbosh",
	}
	result, err := stg.Commander().Create(commander)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestGetByIdCommander(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	var id pb.ById
	id.Id = "17f88543-39e3-43c2-a8e7-e7ce0c245db6"

	commander, err := stg.Commander().GetById(&id)

	assert.NoError(t, err)
	assert.NotNil(t, commander)
}

func TestGetAllCommanders(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}
	commanders, err := stg.Commander().GetAll(&pb.CommanderReq{})
	assert.NoError(t, err)
	assert.NotNil(t, commanders)
}

func TestUpdateCommander(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	commander := &pb.Commander{
		Id:          "17f88543-39e3-43c2-a8e7-e7ce0c245db6",
		Name:        "Jane Doe",
		Email:       "jane.doe@example.com",
		DateOfBirth: "1985-02-15",
		PhoneNumber: "+9876543210",
		Position:    "mingbosh",
	}
	result, err := stg.Commander().Update(commander)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestDeleteCommander(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	var id pb.ById
	id.Id = "80e07ab6-708c-4534-a3b5-9d3e643e78cd"

	result, err := stg.Commander().Delete(&id)

	assert.NoError(t, err)
	assert.Nil(t, result)
}
