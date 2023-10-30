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
	AllocRoleForAdmin(c context.Context, data *model.UmsRoleRelationInReq) error
}

// AllocRoleForAdmin 为后台用户分配角色
func (rs *umsRoleService) AllocRoleForAdmin(c context.Context, data *model.UmsRoleRelationInReq) error {
	// 1. 根据角色ID去查找
	_, exists, err := rs.repo.SelectById(c, data.RoleId)
	if err != nil {
		return e.ErrInternalServer().WithMsg("分配角色失败, 请稍后再试~").WithErr(err)
	}
	if !exists {
		return e.ErrNotFound().WithMsg(fmt.Sprintf("roleId: %v, 角色不存在", data.RoleId))
	}
	panic("unimplemented")
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
	if err := rs.repo.RemoveByRoleId(c, reqData.RoleId); err != nil {
		return e.ErrInternalServer().WithMsg("分配菜单失败, 请稍后再试").WithErr(err)
	}

	list := make([]*model.UmsRoleMenuRelation, len(reqData.MenuIds))
	for i, v := range reqData.MenuIds {
		list[i] = new(model.UmsRoleMenuRelation)
		list[i].RoleId = reqData.RoleId
		list[i].MenuId = v
	}

	if err := rs.repo.BatchInsert(c, list); err != nil {
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
	_, exists, err := rs.repo.SelectByRoleName(c, reqData.Name)
	if err != nil {
		return e.ErrInternalServer().WithMsg("创建角色失败, 请稍后再试")
	}
	if exists {
		return e.ErrBadRequest().WithMsg(fmt.Sprintf("roleName: %v, 角色名称以存在", reqData.Name))
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
