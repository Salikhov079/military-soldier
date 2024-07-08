package service

import (
	"context"
	"log"

	pb "github.com/Salikhov079/military/genprotos/soldiers"
	s "github.com/Salikhov079/military/storage"
)

type DepartmentService struct {
	stg s.InitRoot
	pb.UnimplementedDepartmentServiceServer
}

func NewDepartmentService(stg s.InitRoot) *DepartmentService {
	return &DepartmentService{stg: stg}
}

func (c *DepartmentService) Create(ctx context.Context, department *pb.Department) (*pb.Void, error) {
	void, err := c.stg.Department().Create(department)
	if err != nil {
		log.Print(err)
	}
	return void, err
}

func (c *DepartmentService) Update(ctx context.Context, department *pb.Department) (*pb.Void, error) {
	void, err := c.stg.Department().Update(department)
	if err != nil {
		log.Print(err)
	}
	
	return void, err
}

func (c *DepartmentService) Delete(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	void, err := c.stg.Department().Delete(id)
	if err != nil {
		log.Print(err)
	}

	return void, err
}

func (c *DepartmentService) GetById(ctx context.Context, id *pb.ById) (*pb.Department, error) {
	department, err := c.stg.Department().GetById(id)
	if err != nil {
		log.Print(err)
	}
	
	return department, err
}

func (c *DepartmentService) GetAll(ctx context.Context, filter *pb.Department) (*pb.AllDepartments, error) {
	departments, err := c.stg.Department().GetAll(filter)
	if err != nil {
		log.Print(err)
	}

	return departments, err
}
