package postgres

import (
	"database/sql"

	cp "company/genprotos"
)

type PositionRepo struct {
	db *sql.DB
}

func NewPositionRepo(db *sql.DB) *PositionRepo {
	return &PositionRepo{db: db}
}

func(p *PositionRepo) Create(req *cp.PositionCreateReq) (*cp.PositionRes, error){
	return nil, nil
}

func(p *PositionRepo) GetById(id *cp.Byid) (*cp.PositionGetByIdRes, error){
	return nil, nil
}

func(p *PositionRepo) GetAll(req *cp.PositionGetAllReq) (*cp.PositionGetAllRes, error){
	return nil, nil
}

func(p *PositionRepo) Update(req *cp.PositionUpdateReq) (*cp.PositionRes, error){
	return nil, nil
}

func(p *PositionRepo) Delete(id *cp.Byid) (*cp.Void, error){
	return nil, nil
}
