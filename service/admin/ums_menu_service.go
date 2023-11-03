package adminsrv

import (
	"context"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/e"
	adminrepo "github.com/kalougata/mall/repo/admin"
)

type umsMenuService struct {
	repo adminrepo.UmsMenuRepo
}

type UmsMenuService interface {
	AddMenu(c context.Context, reqData *model.UmsMenuInReq) error
	TreeList(c context.Context) (list []*model.UmsMenuNode, err error)
}

// TreeList 获取树形菜单列表
func (ms *umsMenuService) TreeList(c context.Context) (list []*model.UmsMenuNode, err error) {
	var menuList []*model.UmsMenu
	menuList, err = ms.repo.SelectList(c)
	if err != nil {
		return nil, err
	}
	var treeList []*model.UmsMenuNode
	for _, menu := range menuList {
		if menu.ParentID == 0 {
			node := convertMenuNode(menu, menuList)
			treeList = append(treeList, node)
		}
	}

	return treeList, nil
}

// AddMenu 添加分类
func (ms *umsMenuService) AddMenu(c context.Context, reqData *model.UmsMenuInReq) error {
	if count, err := ms.repo.Create(c, reqData.ToModel()); err != nil || count <= 0 {
		return e.ErrInternalServer().WithMsg("添加分类失败, 请稍后再试")
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
	var children []*model.UmsMenuNode
	for _, subMenu := range list {
		if subMenu.ParentID == menu.ID {
			child := convertMenuNode(subMenu, list)
			children = append(children, child)
		}
	}
	node.Children = children

	return node
}

func NewUmsMenuService(repo adminrepo.UmsMenuRepo) UmsMenuService {
	return &umsMenuService{repo}
}
