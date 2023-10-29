package adminrepo

import (
	"context"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/data"
)

type umsMenuRepo struct {
	data *data.Data
}

type UmsMenuRepo interface {
	Create(c context.Context, menu *model.UmsMenu) error
}

// Create 创建菜单
func (repo *umsMenuRepo) Create(c context.Context, menu *model.UmsMenu) error {
	if _, err := repo.data.DB.Context(c).Insert(menu); err != nil {
		return err
	}

	return nil
}

func NewUmsMenuRepo(data *data.Data) UmsMenuRepo {
	return &umsMenuRepo{data}
}
