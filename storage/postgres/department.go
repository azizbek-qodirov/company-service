package postgres

import (
	"database/sql"

	cp "company/genprotos"
)

type DepartmentRepo struct {
	db *sql.DB
}

func NewDepartmentRepo(db *sql.DB) *DepartmentRepo {
	return &DepartmentRepo{db: db}
}

func (d *DepartmentRepo) Create(req *cp.DepartmentCreateReq) (*cp.DepartmentRes, error) {
	return nil, nil
}

func (d *DepartmentRepo) GetById(id *cp.Byid) (*cp.DepartmentGetByIdRes, error) {
	return nil, nil
}

func (d *DepartmentRepo) GetAll(req *cp.DepartmentGetAllReq) (*cp.DepartmentGetAllRes, error) {
	return nil, nil
}

func (d *DepartmentRepo) Update(req *cp.DepartmentUpdateReq) (*cp.DepartmentRes, error) {
	return nil, nil
}

func (d *DepartmentRepo) Delete(id *cp.Byid) (*cp.Void, error) {
	return nil, nil
}