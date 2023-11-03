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
	Create(c context.Context, category *model.UmsResourceCategory) (int64, error)
}

// Create 创建资源分类
func (repo *umsResourceCategoryRepo) Create(c context.Context, category *model.UmsResourceCategory) (int64, error) {
	return repo.data.DB.Context(c).Table(&model.UmsResourceCategory{}).Insert(category)
}

func NewUmsResourceCategoryRepo(data *data.Data) UmsResourceCategoryRepo {
	return &umsResourceCategoryRepo{data}
}
