package adminsrv

import adminrepo "github.com/kalougata/mall/repo/admin"

type umsAdminService struct {
	repo adminrepo.UmsAdminRepo
}

type UmsAdminService interface {
}

func NewUmsAdminService(repo adminrepo.UmsAdminRepo) UmsAdminService {
	return &umsAdminService{repo}
}
