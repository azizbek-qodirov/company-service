package service

import (
	cp "company/genprotos"
	st "company/storage/postgres"
	"context"
)

type ResumeService struct {
	storage st.Storage
	cp.UnimplementedResumeServiceServer
}

func NewResumeService(storage *st.Storage) *ResumeService {
	return &ResumeService{
		storage: *storage,
	}
}

func (s *ResumeService) Create(ctx context.Context, req *cp.ResumeCreateReq) (*cp.ResumeRes, error) {
	return nil, nil
}

func (s *ResumeService) GetById(ctx context.Context, id *cp.Byid) (*cp.ResumeGetByIdRes, error) {
	return nil, nil
}

func (s *ResumeService) GetAll(ctx context.Context, req *cp.ResumeGetAllReq) (*cp.ResumeGetAllRes, error) {
	return nil, nil
}

func (s *ResumeService) Update(ctx context.Context, req *cp.ResumeUpdateReq) (*cp.ResumeRes, error) {
	return nil, nil
}

func (s *ResumeService) Delete(ctx context.Context, id *cp.Byid) (*cp.Void, error) {
	return nil, nil
}
