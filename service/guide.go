package service

import (
	"context"

	pb "github.com/Azizbek-Qodirov/hr_platform_evaluation_service/genprotos"
	s "github.com/Azizbek-Qodirov/hr_platform_evaluation_service/storage"
)

type GuideService struct {
	stg s.InitRoot
	pb.UnimplementedGuideServiceServer
}

func NewGuideService(stg s.InitRoot) *GuideService {
	return &GuideService{stg: stg}
}

func (s *EvaluationService) CreateGuide(ctx context.Context, req *pb.CreateGuideRequest) (*pb.Void, error) {
	return s.stg.Guide().CreateGuide(req)
}

func (s *EvaluationService) GetGuide(ctx context.Context, req *pb.GetGuideRequest) (*pb.GuideResponse, error) {
	return s.stg.Guide().GetGuide(req)
}

func (s *EvaluationService) UpdateGuide(ctx context.Context, req *pb.UpdateGuideRequest) (*pb.Void, error) {
	return s.stg.Guide().UpdateGuide(req)
}

func (s *EvaluationService) DeleteGuide(ctx context.Context, req *pb.DeleteGuideRequest) (*pb.Void, error) {
	return s.stg.Guide().DeleteGuide(req)
}

func (s *EvaluationService) ListAllGuides(ctx context.Context, req *pb.ListAllGuidesRequest) (*pb.ListAllGuidesResponse, error) {
	return s.stg.Guide().ListAllGuides(req)
}

func (s *EvaluationService) SearchGuides(ctx context.Context, req *pb.SearchGuidesRequest) (*pb.ListAllGuidesResponse, error) {
	return s.stg.Guide().SearchGuides(req)
}
