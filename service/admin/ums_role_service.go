package adminsrv

import (
	"context"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/e"
	adminrepo "github.com/kalougata/mall/repo/admin"
)

type umsRoleService struct {
	repo adminrepo.UmsRoleRepo
}

type UmsRoleService interface {
	AddRole(c context.Context, reqData *model.AddUmsRoleReq) error
	UpdateRole(c context.Context, reqData *model.UpdateUmsRoleReq) error
	BatchDeleteRole(c context.Context, ids []string) error
	ListRole(c context.Context) (list []*model.UmsRole, err error)
}

// ListRole implements UmsRoleService.
func (rs *umsRoleService) ListRole(c context.Context) (list []*model.UmsRole, err error) {
	return rs.repo.SelectList(c)
}

// UpdateRole implements UmsRoleService.
func (rs *umsRoleService) UpdateRole(c context.Context, reqData *model.UpdateUmsRoleReq) error {
	if err := rs.repo.Update(c, reqData.ToModel()); err != nil {
		return e.ErrInternalServer().WithMsg("更新角色失败, 请稍后再试")
	}

	return nil
}

// AddRole 添加一个角色
func (rs *umsRoleService) AddRole(c context.Context, reqData *model.AddUmsRoleReq) error {
	_, exists, err := rs.repo.SelectByRoleName(c, reqData.Name)
	if err != nil {
		return e.ErrInternalServer().WithMsg("创建角色失败, 请稍后再试")
	}
	if exists {
		return e.ErrBadRequest().WithMsg("角色名称已存在")
	}

	return rs.repo.Save(c, reqData.ToModel())
}

// BatchUpdateRole 删除角色, 支持批量删除
func (rs *umsRoleService) BatchDeleteRole(c context.Context, ids []string) error {
	if err := rs.repo.Delete(c, ids); err != nil {
		return e.ErrInternalServer().WithMsg("删除失败, 请稍后再试").WithErr(err)
	}

	return nil
}

func NewUmsRoleService(repo adminrepo.UmsRoleRepo) UmsRoleService {
	return &umsRoleService{repo}
}
