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
	SelectByRoleName(c context.Context, roleName string) (result *model.UmsRole, exists bool, err error)
	SelectList(c context.Context, ids []string) ([]*model.UmsRole, error)
}

// Save 创建一个角色
func (repo *umsRoleRepo) Save(c context.Context, role *model.UmsRole) error {
	if count, err := repo.data.DB.Context(c).Insert(role); err != nil && count <= 0 {
		return err
	}

	return nil
}

// Update implements UmsRoleRepo.
func (repo *umsRoleRepo) Delete(c context.Context, ids []string) error {
	if count, err := repo.data.DB.Context(c).In("id", ids).Delete(&model.UmsRole{}); err != nil && count <= 0 {
		return err
	}

	return nil
}

// SelectList implements UmsRoleRepo.
func (repo *umsRoleRepo) SelectList(c context.Context, ids []string) ([]*model.UmsRole, error) {
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
