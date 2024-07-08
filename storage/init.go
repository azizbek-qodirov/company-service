package storage

import (
	pb "github.com/Azizbek-Qodirov/hr_platform_evaluation_service/genprotos"
)

type InitRoot interface {
	Evaluation() EvaluationStorage
	Guide() GuideStorage
	Notification() NotificationStorage
}

type EvaluationStorage interface {
	GetAll(req *pb.EvaluationGetAllReq) (*pb.EvaluationGetAllRes, error)
	Delete(req *pb.Byid) (*pb.Void, error)
	Update(req *pb.EvaluationUpdate) (*pb.Void, error)
	Get(req *pb.Byid) (*pb.EvaluationUpdate, error)
	Create(req *pb.EvaluationCreate) (*pb.Void, error)
}


type GuideStorage interface {
	CreateGuide(req *pb.CreateGuideRequest) (*pb.Void, error)
	GetGuide(req *pb.GetGuideRequest) (*pb.GuideResponse, error)
	UpdateGuide(req *pb.UpdateGuideRequest) (*pb.Void, error)
	DeleteGuide(req *pb.DeleteGuideRequest) (*pb.Void, error)
	ListAllGuides(req *pb.ListAllGuidesRequest) (*pb.ListAllGuidesResponse, error)
	SearchGuides(req *pb.SearchGuidesRequest) (*pb.ListAllGuidesResponse, error)
}


type NotificationStorage interface {
	Create(notification *pb.CreateNotification) (*pb.Void, error)
	GetByUserId(request *pb.GetAllRequest) (*pb.GetAllResponse, error) 
	ReadeAll(request *pb.ReadeAllRequest) (*pb.Void, error)
	SendAll(request *pb.SendByCompanyidToUsers) (*pb.Void, error)
	SendAllUsers(request *pb.SendByCompanyidToUsers) (*pb.Void, error)
}
