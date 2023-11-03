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
	Create(c context.Context, resource *model.UmsResource) (int64, error)
}

// Create 创建资源
func (repo *umsResourceRepo) Create(c context.Context, resource *model.UmsResource) (int64, error) {
	return repo.data.DB.Context(c).Table(&model.UmsResource{}).Insert(resource)
}

func NewUmsResourceRepo(data *data.Data) UmsResourceRepo {
	return &umsResourceRepo{data}
}
