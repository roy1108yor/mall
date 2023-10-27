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
}

// Save implements UmsRoleRepo.
func (repo *umsRoleRepo) Save(c context.Context, role *model.UmsRole) error {
	if count, err := repo.data.DB.Context(c).Insert(role); err != nil && count <= 0 {
		return e.ErrInternalServer().WithErr(err)
	}

	return nil
}

func NewUmsRoleRepo(data *data.Data) UmsRoleRepo {
	return &umsRoleRepo{data}
}
