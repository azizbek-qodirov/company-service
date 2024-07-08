package postgres

import (
	"database/sql"

	cp "company/genprotos"
)

type CompanyRepo struct {
	db *sql.DB
}

func NewCompanyRepo(db *sql.DB) *CompanyRepo {
	return &CompanyRepo{db: db}
}

func (c *CompanyRepo) Create(req *cp.CompanyCreateReq) (*cp.CompanyRes, error) {
	return nil, nil
}

func (c *CompanyRepo) GetById(id *cp.Byid) (*cp.CompanyGetByIdRes, error) {
	return nil, nil
}

func (c *CompanyRepo) GetAll(req *cp.CompanyGetAllReq) (*cp.CompanyGetAllRes, error) {
	return nil, nil
}

func (c *CompanyRepo) Update(req *cp.CompanyUpdateReq) (*cp.CompanyRes, error) {
	return nil, nil
}

func (c *CompanyRepo) Delete(id *cp.Byid) (*cp.Void, error) {
	return nil, nil
}
