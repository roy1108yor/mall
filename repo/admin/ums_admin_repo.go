package adminrepo

import (
	"context"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/data"
)

type umsAdminRepo struct {
	*data.Data
}

// Save implements UmsAdminRepo.
func (repo *umsAdminRepo) Save(c context.Context, admin *model.UmsAdmin) error {
	if _, err := repo.Data.DB.Context(c).Insert(admin); err != nil {
		return err
	}

	return nil
}

type UmsAdminRepo interface {
	Save(c context.Context, admin *model.UmsAdmin) error
}

func NewUmsAdminRepo(data *data.Data) UmsAdminRepo {
	return &umsAdminRepo{Data: data}
}
