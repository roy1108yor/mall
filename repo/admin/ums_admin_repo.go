package adminrepo

import (
	"context"
	"fmt"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/data"
	"github.com/kalougata/mall/pkg/e"
)

type umsAdminRepo struct {
	*data.Data
}

type UmsAdminRepo interface {
	Save(c context.Context, admin *model.UmsAdmin) error
	SelectByUserName(c context.Context, userName string) (result *model.UmsAdmin, exists bool, err error)
}

// Save implements UmsAdminRepo.
func (repo *umsAdminRepo) Save(c context.Context, admin *model.UmsAdmin) error {
	if _, err := repo.Data.DB.Context(c).Insert(admin); err != nil {
		fmt.Println(err)
		return e.ErrInternalServer().WithErr(err)
	}

	return nil
}

// SelectByUserName implements UmsAdminRepo.
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
