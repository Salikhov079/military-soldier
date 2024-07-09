package main

import (
	"log"
	"net"

	pb "github.com/Salikhov079/military/genprotos/soldiers"
	"github.com/Salikhov079/military/service"
	postgres "github.com/Salikhov079/military/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	liss, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterGroupServiceServer(s, service.NewGroupService(db))
	pb.RegisterSoldierServiceServer(s, service.NewSoldierService(db))
	pb.RegisterCommanderServiceServer(s, service.NewCommanderService(db))
	pb.RegisterDepartmentServiceServer(s, service.NewDepartmentService(db))
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
