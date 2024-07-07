package postgres

import (
	"log"
	"testing"

	pb "github.com/Salikhov079/military/genprotos/soldiers"
	"github.com/Salikhov079/military/storage/postgres"

	"github.com/stretchr/testify/assert"
)

func TestCreateGroup(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	group := &pb.GroupReq{
		Name:         "Alpha Group",
		DepartmentId: "359bbc19-d10c-48f8-89be-e8c83d3f2b7a",
		CommandersId: "17f88543-39e3-43c2-a8e7-e7ce0c245db6",
	}
	result, err := stg.Group().Create(group)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestGetByIdGroup(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	var id pb.ById
	id.Id = "b015266e-87ef-4cba-aa1c-b70a8380415d"

	group, err := stg.Group().GetById(&id)

	assert.NoError(t, err)
	assert.NotNil(t, group)
}

func TestGetAllGroups(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}
	groups, err := stg.Group().GetAll(&pb.GroupReq{})
	assert.NoError(t, err)
	assert.NotNil(t, groups)
}

func TestUpdateGroup(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	group := &pb.Group{
		Id:           "b015266e-87ef-4cba-aa1c-b70a8380415d",
		Name:         "Updated Alpha Group",
		Department: &pb.Department{Id: "359bbc19-d10c-48f8-89be-e8c83d3f2b7a"},
		Commanders: &pb.Commander{Id: "17f88543-39e3-43c2-a8e7-e7ce0c245db6"},
	}
	result, err := stg.Group().Update(group)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestDeleteGroup(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	var id pb.ById
	id.Id = "80e07ab6-708c-4534-a3b5-9d3e643e78cd"

	result, err := stg.Group().Delete(&id)

	assert.NoError(t, err)
	assert.Nil(t, result)
}
