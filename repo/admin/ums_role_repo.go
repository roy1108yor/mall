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
	Save(c context.Context, role *model.UmsRole) error
	Delete(c context.Context, ids []string) error
	Update(c context.Context, role *model.UmsRole) error
	SelectByRoleName(c context.Context, roleName string) (result *model.UmsRole, exists bool, err error)
	SelectList(c context.Context, ids ...string) ([]*model.UmsRole, error)
	SelectById(c context.Context, id uint) (result *model.UmsRole, exists bool, err error)
	RemoveByRoleId(c context.Context, id uint) error
	BatchInsert(c context.Context, list []*model.UmsRoleMenuRelation) error
	InsertRoleRelationForAdmin(c context.Context, relation *model.UmsRoleRelation) error
}

// InsertRoleRelationForAdmin 为后台用户分配角色
func (repo *umsRoleRepo) InsertRoleRelationForAdmin(c context.Context, relation *model.UmsRoleRelation) error {
	// 1. 先查找用户ID是否存在
	panic("unimplemented")
}

// BatchInsert implements UmsRoleRepo.
func (repo *umsRoleRepo) BatchInsert(c context.Context, list []*model.UmsRoleMenuRelation) error {
	if _, err := repo.data.DB.Context(c).Table(&model.UmsRoleMenuRelation{}).InsertMulti(list); err != nil {
		return err
	}

	return nil
}

// RemoveById implements UmsRoleRepo.
func (repo *umsRoleRepo) RemoveByRoleId(c context.Context, roleId uint) error {
	if _, err := repo.data.DB.Context(c).Table(&model.UmsRoleMenuRelation{}).Where("role_id = ?", roleId).Delete(&model.UmsRoleMenuRelation{}); err != nil {
		return err
	}

	return nil
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
func (repo *umsRoleRepo) Save(c context.Context, role *model.UmsRole) error {
	if count, err := repo.data.DB.Context(c).Insert(role); err != nil && count <= 0 {
		return err
	}

	return nil
}

// Delete 批量删除角色
func (repo *umsRoleRepo) Delete(c context.Context, ids []string) error {
	if count, err := repo.data.DB.Context(c).In("id", ids).Delete(&model.UmsRole{}); err != nil && count <= 0 {
		return err
	}

	return nil
}

// SelectList implements UmsRoleRepo.
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
