package service

import (
	"context"
	"log"

	pb "github.com/Salikhov079/military/genprotos/soldiers"
	s "github.com/Salikhov079/military/storage"
)

type CommanderService struct {
	stg s.InitRoot
	pb.UnimplementedCommanderServiceServer
}

func NewCommanderService(stg s.InitRoot) *CommanderService {
	return &CommanderService{stg: stg}
}

func (c *CommanderService) Create(ctx context.Context, commander *pb.CommanderReq) (*pb.Void, error) {
	void, err := c.stg.Commander().Create(commander)
	if err != nil {
		log.Print(err)
	}
	return void, err
}

func (c *CommanderService) Update(ctx context.Context, commander *pb.Commander) (*pb.Void, error) {
	void, err := c.stg.Commander().Update(commander)
	if err != nil {
		log.Print(err)
	}
	
	return void, err
}

func (c *CommanderService) Delete(ctx context.Context, id *pb.ById) (*pb.Void, error) {
	void, err := c.stg.Commander().Delete(id)
	if err != nil {
		log.Print(err)
	}

	return void, err
}

func (c *CommanderService) GetById(ctx context.Context, id *pb.ById) (*pb.Commander, error) {
	commander, err := c.stg.Commander().GetById(id)
	if err != nil {
		log.Print(err)
	}
	
	return commander, err
}

func (c *CommanderService) GetAll(ctx context.Context, filter *pb.CommanderReq) (*pb.AllCommanders, error) {
	commanders, err := c.stg.Commander().GetAll(filter)
	if err != nil {
		log.Print(err)
	}

	return commanders, err
}
