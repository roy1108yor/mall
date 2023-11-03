package adminrepo

import (
	"context"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/data"
	"github.com/kalougata/mall/pkg/e"
)

type umsRoleRepo struct {
	data *data.Data
}

type UmsRoleRepo interface {
	Create(c context.Context, role *model.UmsRole) (int64, error)
	Delete(c context.Context, ids []string) error
	Update(c context.Context, role *model.UmsRole) error
	SelectByRoleName(c context.Context, roleName string) (result *model.UmsRole, exists bool, err error)
	SelectList(c context.Context, ids ...string) ([]*model.UmsRole, error)
	SelectById(c context.Context, id uint) (result *model.UmsRole, exists bool, err error)
	RemoveRoleMenuRelationByRoleId(c context.Context, id uint) (int64, error)
	BatchInsertRoleMenuRelationForRole(c context.Context, list []*model.UmsRoleMenuRelation) (int64, error)
	BatchInsertRoleRelationForAdmin(c context.Context, list []*model.UmsRoleRelation) (int64, error)
}

func (repo *umsRoleRepo) BatchInsertRoleRelationForAdmin(c context.Context, list []*model.UmsRoleRelation) (int64, error) {
	return repo.data.DB.Context(c).Table(&model.UmsRoleRelation{}).InsertMulti(list)
}

func (repo *umsRoleRepo) BatchInsertRoleMenuRelationForRole(c context.Context, list []*model.UmsRoleMenuRelation) (int64, error) {
	return repo.data.DB.Context(c).Table(&model.UmsRoleMenuRelation{}).InsertMulti(list)
}

func (repo *umsRoleRepo) RemoveRoleMenuRelationByRoleId(c context.Context, roleId uint) (int64, error) {
	return repo.data.DB.Context(c).Table(&model.UmsRoleMenuRelation{}).Where("role_id = ?", roleId).Delete(&model.UmsRoleMenuRelation{})
}

// SelectById 根据角色ID查找
func (repo *umsRoleRepo) SelectById(c context.Context, id uint) (result *model.UmsRole, exists bool, err error) {
	result = &model.UmsRole{}
	exists, err = repo.data.DB.Context(c).Where("id = ?", id).Get(result)
	if err != nil {
		err = e.ErrInternalServer().WithErr(err)
	}

	return
}

// Update 更新一个角色
func (repo *umsRoleRepo) Update(c context.Context, role *model.UmsRole) error {
	if count, err := repo.data.DB.Context(c).AllCols().ID(role.ID).Update(role); err != nil && count <= 0 {
		return err
	}

	return nil
}

// Save 创建一个角色
func (repo *umsRoleRepo) Create(c context.Context, role *model.UmsRole) (int64, error) {
	return repo.data.DB.Context(c).Insert(role)
}

// Delete 批量删除角色
func (repo *umsRoleRepo) Delete(c context.Context, ids []string) error {
	if count, err := repo.data.DB.Context(c).In("id", ids).Delete(&model.UmsRole{}); err != nil && count <= 0 {
		return err
	}

	return nil
}

func (repo *umsRoleRepo) SelectList(c context.Context, ids ...string) ([]*model.UmsRole, error) {
	list := make([]*model.UmsRole, 0)
	err := repo.data.DB.Context(c).In("id", ids).Find(&list)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// SelectByRoleName 根据角色名称查找角色
func (repo *umsRoleRepo) SelectByRoleName(c context.Context, roleName string) (result *model.UmsRole, exists bool, err error) {
	result = &model.UmsRole{}
	exists, err = repo.data.DB.Context(c).Where("name = ?", roleName).Get(result)
	if err != nil {
		err = e.ErrInternalServer().WithErr(err)
	}

	return
}

func NewUmsRoleRepo(data *data.Data) UmsRoleRepo {
	return &umsRoleRepo{data}
}
