package adminsrv

import (
	"context"
	"fmt"

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
	AllocMenuForRole(c context.Context, reqData *model.AllocMenuForRoleReq) error
}

// AllocMenuForRole 为角色分配菜单
func (rs *umsRoleService) AllocMenuForRole(c context.Context, reqData *model.AllocMenuForRoleReq) error {
	_, exists, err := rs.repo.SelectById(c, reqData.RoleId)
	if err != nil {
		return err
	}
	if !exists {
		return e.ErrNotFound().WithMsg(fmt.Sprintf("roleId: %v, 角色不存在", reqData.RoleId))
	}

	// 删除角色原有的关系
	if count, err := rs.repo.RemoveRoleMenuRelationByRoleId(c, reqData.RoleId); err != nil || count <= 0 {
		return e.ErrInternalServer().WithMsg("分配菜单失败, 请稍后再试").WithErr(err)
	}

	list := make([]*model.UmsRoleMenuRelation, len(reqData.MenuIds))
	for i, v := range reqData.MenuIds {
		list[i] = new(model.UmsRoleMenuRelation)
		list[i].RoleId = reqData.RoleId
		list[i].MenuId = v
	}

	if _, err := rs.repo.BatchInsertRoleMenuRelationForRole(c, list); err != nil {
		return e.ErrInternalServer().WithMsg("分配菜单失败, 请稍后再试").WithErr(err)
	}

	return nil
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
	_, exists, _ := rs.repo.SelectByRoleName(c, reqData.Name)
	if exists {
		return e.ErrBadRequest().WithMsg(fmt.Sprintf("roleName: %v, 角色名称以存在", reqData.Name))
	}
	if _, err := rs.repo.Create(c, reqData.ToModel()); err != nil {
		return e.ErrInternalServer().WithMsg("创建角色失败, 请稍后再试~")
	}

	return nil
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
