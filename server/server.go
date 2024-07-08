package main

import (
	"log"
	"net"
	pb "github.com/Azizbek-Qodirov/hr_platform_evaluation_service/genprotos"
	"github.com/Azizbek-Qodirov/hr_platform_evaluation_service/service"
	postgres "github.com/Azizbek-Qodirov/hr_platform_evaluation_service/storage/postgres"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
	liss, err := net.Listen("tcp", ":8086")
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterServiceNotificationServer(s, service.NewNotificationService(db))
	pb.RegisterEvaulationServiceServer(s, service.NewEvaluationService(db))
	pb.RegisterGuideServiceServer(s, service.NewGuideService(db))
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	
}
