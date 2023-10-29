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
	TreeList(c context.Context) ([]*model.UmsMenuNode, error)
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

// TreeList implements UmsMenuRepo.
func (repo *umsMenuRepo) TreeList(c context.Context) ([]*model.UmsMenuNode, error) {
	menuList := make([]*model.UmsMenu, 0)
	if err := repo.data.DB.Context(c).Table(&model.UmsMenu{}).Find(&menuList); err != nil {
		return nil, e.ErrInternalServer()
	}
	result := []*model.UmsMenuNode{}
	for _, menu := range menuList {
		if menu.ParentID == 0 {
			node := convertMenuNode(menu, menuList)
			result = append(result, node)
		}
	}

	return result, nil
}

// Create 创建菜单
func (repo *umsMenuRepo) Create(c context.Context, menu *model.UmsMenu) error {
	if _, err := repo.data.DB.Context(c).Insert(menu); err != nil {
		return err
	}

	return nil
}

func convertMenuNode(menu *model.UmsMenu, list []*model.UmsMenu) *model.UmsMenuNode {
	node := &model.UmsMenuNode{
		ParentId: menu.ParentID,
		Name:     menu.Name,
		Icon:     menu.Icon,
		Sort:     menu.Sort,
		Hidden:   menu.Hidden,
	}
	children := []*model.UmsMenuNode{}
	for _, subMenu := range list {
		if subMenu.ParentID == menu.ID {
			child := convertMenuNode(subMenu, list)
			children = append(children, child)
		}
	}
	node.Children = children

	return node
}

func NewUmsMenuRepo(data *data.Data) UmsMenuRepo {
	return &umsMenuRepo{data}
}
