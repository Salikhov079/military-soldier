package postgres

import (
	"log"
	"testing"

	pb "github.com/Salikhov079/military/genprotos/soldiers"
	"github.com/Salikhov079/military/storage/postgres"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func TestCreateSoldier(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	soldier := &pb.SoldierReq{
		Name:        "Your_Name",
		Email:       "e@example.com",
		DateOfBirth: "2000-01-01T00:00:00Z",
		PhoneNumber: "+998901234567",
		GroupId:     "b015266e-87ef-4cba-aa1c-b70a8380415d",
		JoinDate:    "2024-01-01T00:00:00Z",
	}
	result, err := stg.Soldier().Create(soldier)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestGetByIdSoldier(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	var id pb.ById
	id.Id = "097ff004-a6ee-4eee-92bd-d3e0834ba10e"

	soldier, err := stg.Soldier().GetById(&id)

	assert.NoError(t, err)
	assert.NotNil(t, soldier)
}

func TestGetAllSoldiers(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}
	soldiers, err := stg.Soldier().GetAll(&pb.SoldierReq{})
	assert.NoError(t, err)
	assert.NotNil(t, soldiers)
}

func TestUpdateSoldier(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	soldier := &pb.Soldier{
		Id:          "097ff004-a6ee-4eee-92bd-d3e0834ba10e",
		Name:        "New_Soldier_Name",
		Email:       "new.e@example.com",
		DateOfBirth: "2000-01-01T00:00:00Z",
		PhoneNumber: "+998901234567",
		Group: &pb.Group{
			Id: "b015266e-87ef-4cba-aa1c-b70a8380415d",
		},
		JoinDate: "2024-01-01T00:00:00Z",
	}
	result, err := stg.Soldier().Update(soldier)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestDeleteSoldier(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	var id pb.ById
	id.Id = "80e07ab6-708c-4534-a3b5-9d3e643e78cd"

	result, err := stg.Soldier().Delete(&id)

	assert.NoError(t, err)
	assert.Nil(t, result)
}

func TestUseBullet(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	use := &pb.UseB{
		QuantityWeapon:    50,
		QuantityBigWeapon: 5,
		SoldierId:         "097ff004-a6ee-4eee-92bd-d3e0834ba10e",
		Date:              "2024-07-07",
	}
	result, err := stg.Soldier().UseBullet(use)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestUseFuel(t *testing.T) {
	stg, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connecting to db: ", err.Error())
	}

	use := &pb.UseF{
		Diesel:    50,
		Petrol:    50,
		SoldierId: uuid.NewString(),
		Date:      "2024-07-08",
	}
	result, err := stg.Soldier().UseFuel(use)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}