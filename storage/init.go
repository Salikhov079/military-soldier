package postgres

import (
	pb "github.com/Salikhov079/military/genprotos/soldiers"
)

type InitRoot interface {
	Soldier() Soldier
	Group() Group
	Department() Department
	Commander() Commander
}

type Soldier interface {
	Create(soldier *pb.SoldierReq) (*pb.Void, error)
	Update(soldier *pb.Soldier) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.Soldier, error)
	GetAll(filter *pb.SoldierReq) (*pb.AllSoldiers, error)
	UseBullet(use *pb.UseB) (*pb.Void, error)
	UseFuel(use *pb.UseF) (*pb.Void, error)
	GetAllWeaponStatistik(filter *pb.GetSoldierStatistik) (*pb.GetSoldierStatistikRes, error)
	GetAllFuelStatistik(filter *pb.GetSoldierStatistikFuel) (*pb.GetSoldierStatistikFuelRes, error)
}
  

type Group interface {
	Create(group *pb.GroupReq) (*pb.Void, error)
	Update(group *pb.Group) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.Group, error)
	GetAll(filter *pb.GroupReq) (*pb.AllGroups, error)
}

type Department interface {
	Create(department *pb.Department) (*pb.Void, error)
	Update(department *pb.Department) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.Department, error)
	GetAll(filter *pb.Department) (*pb.AllDepartments, error)
}

type Commander interface {
	Create(commander *pb.CommanderReq) (*pb.Void, error)
	Update(commander *pb.Commander) (*pb.Void, error)
	Delete(id *pb.ById) (*pb.Void, error)
	GetById(id *pb.ById) (*pb.Commander, error)
	GetAll(filter *pb.CommanderReq) (*pb.AllCommanders, error)
}
