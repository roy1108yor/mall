package adminrepo

import (
	"context"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/data"
	"github.com/kalougata/mall/pkg/e"
)

type umsMenuRepo struct {
	data *data.Data
}

type UmsMenuRepo interface {
	Create(c context.Context, menu *model.UmsMenu) error
	SelectById(c context.Context, id uint) (menu *model.UmsMenu, exists bool, err error)
	SelectList(c context.Context) (list []*model.UmsMenu, err error)
}

// SelectList 获取菜单列表
func (repo *umsMenuRepo) SelectList(c context.Context) (list []*model.UmsMenu, err error) {
	list = make([]*model.UmsMenu, 0)
	err = repo.data.DB.Context(c).Table(&model.UmsMenu{}).Find(&list)

	return
}

// SelectById implements UmsMenuRepo.
func (repo *umsMenuRepo) SelectById(c context.Context, id uint) (menu *model.UmsMenu, exists bool, err error) {
	menu = &model.UmsMenu{}
	exists, err = repo.data.DB.Context(c).Table(menu).Where("id = ?", id).Get(menu)
	if err != nil {
		err = e.ErrInternalServer()
	}

	return
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
