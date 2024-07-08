package service

import (
	"context"
	"log"

	pb "github.com/Azizbek-Qodirov/hr_platform_evaluation_service/genprotos"
	s "github.com/Azizbek-Qodirov/hr_platform_evaluation_service/storage"
)

type NotificationService struct {
	stg s.InitRoot
	pb.UnimplementedServiceNotificationServer
}

func NewNotificationService(stg s.InitRoot) *NotificationService {
	return &NotificationService{stg: stg}
}

func (s *NotificationService) CreateNotification(ctx context.Context, req *pb.CreateNotification) (*pb.Void, error) {
	resp, err := s.stg.Notification().Create(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *NotificationService) GetNotificationsByUserId(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	resp, err := s.stg.Notification().GetByUserId(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *NotificationService) MarkAllNotificationsAsRead(ctx context.Context, req *pb.ReadeAllRequest) (*pb.Void, error) {
	resp, err := s.stg.Notification().ReadeAll(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *NotificationService) SendNotificationToAll(ctx context.Context, req *pb.SendByCompanyidToUsers) (*pb.Void, error) {
	resp, err := s.stg.Notification().SendAll(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *NotificationService) SendAllUsers(ctx context.Context,request *pb.SendByCompanyidToUsers) (*pb.Void, error){
	res, err:=s.stg.Notification().SendAllUsers(request)
	if err != nil {
		log.Print(err)
	}
	return res, err
}