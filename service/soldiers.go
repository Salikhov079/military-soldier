package service

import (
	"context"
	"log"

	pb "github.com/Salikhov079/military/genprotos/soldiers"
	s "github.com/Salikhov079/military/storage"
)

type SoldierService struct {
	stg s.InitRoot
	pb.UnimplementedSoldierServiceServer
}

func NewSoldierService(stg s.InitRoot) *SoldierService {
	return &SoldierService{stg: stg}
}

func (c *SoldierService) Create(ctx context.Context, soldier *pb.SoldierReq) (*pb.Void, error) {
	void, err := c.stg.Soldier().Create(soldier)
	if err != nil {
		log.Print(err)
	}
	return void, err
}

func (c *SoldierService) Update(ctx context.Context, soldier *pb.Soldier) (*pb.Void, error) {
	void, err := c.stg.Soldier().Update(soldier)
	if err != nil {
		log.Print(err)
	}
	
	return void, err
}

func (c *SoldierService) Delete(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	void, err := c.stg.Soldier().Delete(id)
	if err != nil {
		log.Print(err)
	}

	return void, err
}

func (c *SoldierService) GetById(ctx context.Context, id *pb.ById) (*pb.Soldier, error) {
	soldier, err := c.stg.Soldier().GetById(id)
	if err != nil {
		log.Print(err)
	}
	
	return soldier, err
}

func (c *SoldierService) GetAll(ctx context.Context, filter *pb.SoldierReq) (*pb.AllSoldiers, error) {
	soldiers, err := c.stg.Soldier().GetAll(filter)
	if err != nil {
		log.Print(err)
	}

	return soldiers, err
}

func (c *SoldierService) StatistikWeapons(ctx context.Context, filter *pb.GetSoldierStatistik) (*pb.GetSoldierStatistikRes, error) {
	soldiers, err := c.stg.Soldier().GetAllWeaponStatistik(filter)
	if err != nil {
		log.Print(err)
	}

	return soldiers, err
}



func (c *SoldierService) FuelStatistik(ctx context.Context, filter *pb.GetSoldierStatistikFuel) (*pb.GetSoldierStatistikFuelRes, error) {
	soldiers, err := c.stg.Soldier().GetAllFuelStatistik(filter)
	if err != nil {
		log.Print(err)
	}

	return soldiers, err
}
