package postgres

import (
	"database/sql"

	cp "company/genprotos"
)

type ResumeRepo struct {
	db *sql.DB
}

func NewResumeRepo(db *sql.DB) *ResumeRepo {
	return &ResumeRepo{db: db}
}

func (r *ResumeRepo) Create(req *cp.ResumeCreateReq) (*cp.ResumeRes, error) {
	return nil, nil
}

func (r *ResumeRepo) GetById(id *cp.Byid) (*cp.ResumeGetByIdRes, error) {
	return nil, nil
}

func (r *ResumeRepo) GetAll(req *cp.ResumeGetAllReq) (*cp.ResumeGetAllRes, error) {
	return nil, nil
}

func (r *ResumeRepo) Update(req *cp.ResumeUpdateReq) (*cp.ResumeRes, error) {
	return nil, nil
}

func (r *ResumeRepo) Delete(id *cp.Byid) (*cp.Void, error) {
	return nil, nil
}
