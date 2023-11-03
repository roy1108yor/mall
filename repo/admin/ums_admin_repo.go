package adminrepo

import (
	"context"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/data"
	"github.com/kalougata/mall/pkg/e"
)

type umsAdminRepo struct {
	*data.Data
}

type UmsAdminRepo interface {
	Create(c context.Context, admin *model.UmsAdmin) (int64, error)
	SelectByUserName(c context.Context, userName string) (result *model.UmsAdmin, exists bool, err error)
}

// Create 创建一个用户
func (repo *umsAdminRepo) Create(c context.Context, admin *model.UmsAdmin) (int64, error) {
	return repo.Data.DB.Context(c).Insert(admin)
}

// SelectByUserName 根据用户名查找用户
func (repo *umsAdminRepo) SelectByUserName(c context.Context, userName string) (result *model.UmsAdmin, exists bool, err error) {
	result = &model.UmsAdmin{}
	exists, err = repo.DB.Context(c).Where("user_name = ?", userName).Get(result)
	if err != nil {
		err = e.ErrInternalServer().WithErr(err)
	}

	return
}

func NewUmsAdminRepo(data *data.Data) UmsAdminRepo {
	return &umsAdminRepo{Data: data}
}
