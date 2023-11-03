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
	Create(c context.Context, menu *model.UmsMenu) (int64, error)
	SelectList(c context.Context) (list []*model.UmsMenu, err error)
}

// SelectList 获取菜单列表
func (repo *umsMenuRepo) SelectList(c context.Context) (list []*model.UmsMenu, err error) {
	list = make([]*model.UmsMenu, 0)
	err = repo.data.DB.Context(c).Table(&model.UmsMenu{}).Find(&list)

	return
}

// Create 创建菜单
func (repo *umsMenuRepo) Create(c context.Context, menu *model.UmsMenu) (int64, error) {
	return repo.data.DB.Context(c).Insert(menu)
}

func NewUmsMenuRepo(data *data.Data) UmsMenuRepo {
	return &umsMenuRepo{data}
}
