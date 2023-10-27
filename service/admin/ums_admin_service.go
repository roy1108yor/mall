package adminsrv

import (
	"context"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/e"
	adminrepo "github.com/kalougata/mall/repo/admin"
)

type umsAdminService struct {
	repo adminrepo.UmsAdminRepo
}

type UmsAdminService interface {
	UmsAdminRegister(c context.Context, reqData *model.UmsAdminRegisterReq) error
	UmsAdminLogin(c context.Context, reqData *model.UmsAdminLoginReq) (respData *model.UmsAdminLoginResp, err error)
}

// UmsAdminRegister implements UmsAdminService.
func (us *umsAdminService) UmsAdminRegister(c context.Context, reqData *model.UmsAdminRegisterReq) error {
	_, exists, err := us.repo.SelectByUserName(c, reqData.UserName)
	if err != nil {
		return err
	}
	if exists {
		return e.ErrNotFound().WithMsg("用户名已被注册, 请重新输入")
	}

	admin := &model.UmsAdmin{
		UserName:  reqData.UserName,
		Passwd:    reqData.PassWord,
		RegIpAddr: reqData.RegIpAddr,
	}

	return us.repo.Save(c, admin)
}

// UmsAdminLogin implements UmsAdminService.
func (us *umsAdminService) UmsAdminLogin(c context.Context, reqData *model.UmsAdminLoginReq) (*model.UmsAdminLoginResp, error) {
	_, exists, err := us.repo.SelectByUserName(c, reqData.UserName)
	if !exists {
		return nil, e.ErrNotFound().WithMsg("用户不存在, 可能未注册")
	}
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func NewUmsAdminService(repo adminrepo.UmsAdminRepo) UmsAdminService {
	return &umsAdminService{repo}
}
