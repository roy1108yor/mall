package adminsrv

import (
	"context"

	"github.com/kalougata/mall/model"
	adminrepo "github.com/kalougata/mall/repo/admin"
)

type umsRoleService struct {
	repo adminrepo.UmsRoleRepo
}

type UmsRoleService interface {
	AddRole(c context.Context, reqData *model.UmsRoleReq) error
}

// AddRole implements UmsRoleService.
func (rs *umsRoleService) AddRole(c context.Context, reqData *model.UmsRoleReq) error {
	roleModel := &model.UmsRole{
		Name:        reqData.Name,
		Description: reqData.Description,
		Status:      reqData.Status,
	}

	return rs.repo.Save(c, roleModel)
}

func NewUmsRoleService(repo adminrepo.UmsRoleRepo) UmsRoleService {
	return &umsRoleService{repo}
}
