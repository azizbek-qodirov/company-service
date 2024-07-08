package service

import (
	cp "company/genprotos"
	st "company/storage/postgres"
	"context"
)

type CompanyService struct {
	storage st.Storage
	cp.UnimplementedCompanyServiceServer
}

func NewCompanyService(storage *st.Storage) *CompanyService {
	return &CompanyService{
		storage: *storage,
	}
}

func (s *CompanyService) Create(ctx context.Context, req *cp.CompanyCreateReq) (*cp.CompanyRes, error) {
	return nil, nil
}

func (s *CompanyService) GetById(ctx context.Context, id *cp.Byid) (*cp.CompanyGetByIdRes, error) {
	return nil, nil
}

func (s *CompanyService) GetAll(ctx context.Context, req *cp.CompanyGetAllReq) (*cp.CompanyGetAllRes, error) {
	return nil, nil
}

func (s *CompanyService) Update(ctx context.Context, req *cp.CompanyUpdateReq) (*cp.CompanyRes, error) {
	return nil, nil
}

func (s *CompanyService) Delete(ctx context.Context, id *cp.Byid) (*cp.Void, error) {
	return nil, nil
}
