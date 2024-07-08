package service

import (
	"context"
	"log"

	pb "github.com/Salikhov079/military/genprotos/soldiers"
	s "github.com/Salikhov079/military/storage"
)

type GroupService struct {
	stg s.InitRoot
	pb.UnimplementedGroupServiceServer
}

func NewGroupService(stg s.InitRoot) *GroupService {
	return &GroupService{stg: stg}
}

func (c *GroupService) Create(ctx context.Context, group *pb.GroupReq) (*pb.Void, error) {
	void, err := c.stg.Group().Create(group)
	if err != nil {
		log.Print(err)
	}
	return void, err
}

func (c *GroupService) Update(ctx context.Context, group *pb.Group) (*pb.Void, error) {
	void, err := c.stg.Group().Update(group)
	if err != nil {
		log.Print(err)
	}
	
	return void, err
}

func (c *GroupService) Delete(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	void, err := c.stg.Group().Delete(id)
	if err != nil {
		log.Print(err)
	}

	return void, err
}

func (c *GroupService) GetById(ctx context.Context, id *pb.ById) (*pb.Group, error) {
	group, err := c.stg.Group().GetById(id)
	if err != nil {
		log.Print(err)
	}
	
	return group, err
}

func (c *GroupService) GetAll(ctx context.Context, filter *pb.GroupReq) (*pb.AllGroups, error) {
	groups, err := c.stg.Group().GetAll(filter)
	if err != nil {
		log.Print(err)
	}

	return groups, err
}
