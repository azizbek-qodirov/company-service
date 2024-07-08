package service

import (
	cp "company/genprotos"
	st "company/storage/postgres"
	"context"
)

type DepartmentService struct {
	storage st.Storage
	cp.UnimplementedResumeServiceServer
}

func NewDepartmentService(storage *st.Storage) *DepartmentService {
	return &DepartmentService{
		storage: *storage,
	}
}

func (s *DepartmentService) Create(ctx context.Context, req *cp.DepartmentCreateReq) (*cp.DepartmentRes, error) {
	return nil, nil
}

func (s *DepartmentService) GetById(ctx context.Context, id *cp.Byid) (*cp.DepartmentGetByIdRes, error) {
	return nil, nil
}

func (s *DepartmentService) GetAll(ctx context.Context, req *cp.DepartmentGetAllReq) (*cp.DepartmentGetAllRes, error) {
	return nil, nil
}

func (s *DepartmentService) Update(ctx context.Context, req *cp.DepartmentUpdateReq) (*cp.DepartmentRes, error) {
	return nil, nil
}

func (s *DepartmentService) Delete(ctx context.Context, id *cp.Byid) (*cp.Void, error) {
	return nil, nil
}
