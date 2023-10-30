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
	AddRecourceCategory(c context.Context, data *model.UmsResourceCategoryInReq) error
}

// AddRecourceCategory implements UmsResourceService.
func (rs *umsResourceService) AddRecourceCategory(c context.Context, data *model.UmsResourceCategoryInReq) error {
	if err := rs.resourceCategoryRepo.Create(c, data.ToModel()); err != nil {
		return e.ErrInternalServer().WithErr(err).WithMsg("添加资源分类失败, 请稍后再试~")
	}

	return nil
}

// AddResource implements UmsResourceService.
func (rs *umsResourceService) AddResource(c context.Context, data *model.UmsResourceInReq) error {
	if err := rs.resourceRepo.Create(c, data.ToModel()); err != nil {
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
