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
	ListWithTree(c context.Context) ([]*model.UmsMenuNode, error)
}

// AddMenu 添加分类
func (ms *umsMenuService) AddMenu(c context.Context, reqData *model.UmsMenuInReq) error {
	if err := ms.repo.Create(c, reqData.ToModel()); err != nil {
		return e.ErrInternalServer().WithMsg("添加分类失败, 请稍后再试")
	}

	return nil
}

// ListWithTree implements UmsMenuService.
func (ms *umsMenuService) ListWithTree(c context.Context) ([]*model.UmsMenuNode, error) {
	treeList, err := ms.repo.TreeList(c)
	if err != nil {
		return nil, err
	}

	return treeList, nil
}

func NewUmsMenuService(repo adminrepo.UmsMenuRepo) UmsMenuService {
	return &umsMenuService{repo}
}
