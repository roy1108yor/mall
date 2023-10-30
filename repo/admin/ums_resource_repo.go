package adminrepo

import (
	"context"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/data"
)

type umsResourceRepo struct {
	data *data.Data
}

type UmsResourceRepo interface {
	Create(c context.Context, resource *model.UmsResource) error
}

// Create implements UmsResourceRepo.
func (repo *umsResourceRepo) Create(c context.Context, resource *model.UmsResource) error {
	if _, err := repo.data.DB.Context(c).Table(&model.UmsResource{}).Insert(resource); err != nil {
		return err
	}

	return nil
}

func NewUmsResourceRepo(data *data.Data) UmsResourceRepo {
	return &umsResourceRepo{data}
}
