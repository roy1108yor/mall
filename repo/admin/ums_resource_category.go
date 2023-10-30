package adminrepo

import (
	"context"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/data"
)

type umsResourceCategoryRepo struct {
	data *data.Data
}

type UmsResourceCategoryRepo interface {
	Create(c context.Context, category *model.UmsResourceCategory) error
}

// Create implements UmsResourceCategoryRepo.
func (repo *umsResourceCategoryRepo) Create(c context.Context, category *model.UmsResourceCategory) error {
	if _, err := repo.data.DB.Context(c).Table(&model.UmsResourceCategory{}).Insert(category); err != nil {
		return err
	}

	return nil
}

func NewUmsResourceCategoryRepo(data *data.Data) UmsResourceCategoryRepo {
	return &umsResourceCategoryRepo{data}
}
