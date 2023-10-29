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
	AddMenu(c context.Context, reqData *model.AddUmsMenuReq) error
}

// AddMenu 添加分类
func (ms *umsMenuService) AddMenu(c context.Context, reqData *model.AddUmsMenuReq) error {
	if err := ms.repo.Create(c, reqData.ToModel()); err != nil {
		return e.ErrInternalServer().WithMsg("添加分类失败, 请稍后再试")
	}

	return nil
}

func NewUmsMenuService(repo adminrepo.UmsMenuRepo) UmsMenuService {
	return &umsMenuService{repo}
}
