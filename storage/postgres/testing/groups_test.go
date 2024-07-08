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
		DepartmentId: "68ed89b8-63c3-443d-8a2c-1264a563c905",
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
	id.Id = "6bf3ee89-4bc2-453d-97c6-69a1697ac556"

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
		Id:           "6bf3ee89-4bc2-453d-97c6-69a1697ac556",
		Name:         "Updated Alpha Group",
		Department: &pb.Department{Id: "68ed89b8-63c3-443d-8a2c-1264a563c905"},
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
