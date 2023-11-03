package adminsrv

import (
	"context"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/e"
	adminrepo "github.com/kalougata/mall/repo/admin"
)

type umsResourceService struct {
	resourceRepo         adminrepo.UmsResourceRepo
	resourceCategoryRepo adminrepo.UmsResourceCategoryRepo
}

type UmsResourceService interface {
	AddResource(c context.Context, data *model.UmsResourceInReq) error
	AddResourceCategory(c context.Context, data *model.UmsResourceCategoryInReq) error
}

// AddResourceCategory 添加资源分类
func (rs *umsResourceService) AddResourceCategory(c context.Context, data *model.UmsResourceCategoryInReq) error {
	if count, err := rs.resourceCategoryRepo.Create(c, data.ToModel()); err != nil || count <= 0 {
		return e.ErrInternalServer().WithErr(err).WithMsg("添加资源分类失败, 请稍后再试~")
	}

	return nil
}

// AddResource 添加资源
func (rs *umsResourceService) AddResource(c context.Context, data *model.UmsResourceInReq) error {
	if count, err := rs.resourceRepo.Create(c, data.ToModel()); err != nil || count <= 0 {
		return e.ErrInternalServer().WithErr(err).WithMsg("添加资源失败, 请稍后再试~")
	}

	return nil
}

func NewUmsResourceService(
	resourceRepo adminrepo.UmsResourceRepo,
	resourceCategoryRepo adminrepo.UmsResourceCategoryRepo,
) UmsResourceService {
	return &umsResourceService{resourceRepo, resourceCategoryRepo}
}
