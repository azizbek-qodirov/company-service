package service

import (
	pb "company/genprotos"
	"company/storage"
	"context"
	"log"
)

type EvaluationService struct {
	stg storage.InitRoot
	pb.UnimplementedEvaulationServiceServer
}

func NewEvaluationService(stg storage.InitRoot) *EvaluationService {
	return &EvaluationService{stg: stg}
}

func (s *EvaluationService) Create(ctx context.Context, req *pb.EvaluationCreate) (*pb.Void, error) {
	resp, err := s.stg.Evaluation().Create(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *EvaluationService) Get(ctx context.Context, req *pb.Byid) (*pb.EvaluationUpdate, error) {
	resp, err := s.stg.Evaluation().Get(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *EvaluationService) Update(ctx context.Context, req *pb.EvaluationUpdate) (*pb.Void, error) {
	resp, err := s.stg.Evaluation().Update(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *EvaluationService) Delete(ctx context.Context, req *pb.Byid) (*pb.Void, error) {
	resp, err := s.stg.Evaluation().Delete(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}

func (s *EvaluationService) GetAll(ctx context.Context, req *pb.EvaluationGetAllReq) (*pb.EvaluationGetAllRes, error) {
	resp, err := s.stg.Evaluation().GetAll(req)
	if err != nil {
		log.Print(err)
	}
	return resp, err
}
