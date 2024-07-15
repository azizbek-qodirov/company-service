package storage

import cp "company/genprotos"

type CompanyI interface {
	Create(*cp.CompanyCreateReq) (*cp.CompanyRes, error)
	GetById(*cp.Byid) (*cp.CompanyGetByIdRes, error)
	GetAll(*cp.CompanyGetAllReq) (*cp.CompanyGetAllRes, error)
	Update(*cp.CompanyUpdateReq) (*cp.CompanyRes, error)
	Delete(*cp.Byid) (*cp.Void, error)
}

type DepartmentI interface {
	Create(*cp.DepartmentCreateReq) (*cp.DepartmentRes, error)
	GetById(*cp.Byid) (*cp.DepartmentGetByIdRes, error)
	GetAll(*cp.DepartmentGetAllReq) (*cp.DepartmentGetAllRes, error)
	Update(*cp.DepartmentUpdateReq) (*cp.DepartmentRes, error)
	Delete(*cp.Byid) (*cp.Void, error)
}

type PositionI interface {
	Create(*cp.PositionCreateReq) (*cp.PositionRes, error)
	GetById(*cp.Byid) (*cp.PositionGetByIdRes, error)
	GetAll(*cp.PositionGetAllReq) (*cp.PositionGetAllRes, error)
	Update(*cp.PositionUpdateReq) (*cp.PositionRes, error)
	Delete(*cp.Byid) (*cp.Void, error)
}

type ResumeI interface {
	Create(*cp.ResumeCreateReq) (*cp.ResumeRes, error)
	GetById(*cp.Byid) (*cp.ResumeGetByIdRes, error)
	GetAll(*cp.ResumeGetAllReq) (*cp.ResumeGetAllRes, error)
	Update(*cp.ResumeUpdateReq) (*cp.ResumeRes, error)
	Delete(*cp.Byid) (*cp.Void, error)
}

type InitRoot interface {
	Evaluation() EvaluationStorage
	Guide() GuideStorage
	Notification() NotificationStorage
}

type EvaluationStorage interface {
	GetAll(req *cp.EvaluationGetAllReq) (*cp.EvaluationGetAllRes, error)
	Delete(req *cp.Byid) (*cp.Void, error)
	Update(req *cp.EvaluationUpdate) (*cp.Void, error)
	Get(req *cp.Byid) (*cp.EvaluationUpdate, error)
	Create(req *cp.EvaluationCreate) (*cp.Void, error)
}

type GuideStorage interface {
	CreateGuide(req *cp.CreateGuideRequest) (*cp.Void, error)
	GetGuide(req *cp.GetGuideRequest) (*cp.GuideResponse, error)
	UpdateGuide(req *cp.UpdateGuideRequest) (*cp.Void, error)
	DeleteGuide(req *cp.DeleteGuideRequest) (*cp.Void, error)
	ListAllGuides(req *cp.ListAllGuidesRequest) (*cp.ListAllGuidesResponse, error)
	SearchGuides(req *cp.SearchGuidesRequest) (*cp.ListAllGuidesResponse, error)
}

type NotificationStorage interface {
	Create(notification *cp.CreateNotification) (*cp.Void, error)
	GetByUserId(request *cp.GetAllRequest) (*cp.GetAllResponse, error)
	ReadeAll(request *cp.ReadeAllRequest) (*cp.Void, error)
	SendAll(request *cp.SendByCompanyidToUsers) (*cp.Void, error)
	SendAllUsers(request *cp.SendByCompanyidToUsers) (*cp.Void, error)
}
