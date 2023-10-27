package adminsrv

import (
	"context"
	"time"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/e"
	"github.com/kalougata/mall/pkg/hash"
	"github.com/kalougata/mall/pkg/jwt"
	adminrepo "github.com/kalougata/mall/repo/admin"
	"github.com/spf13/viper"
)

type umsAdminService struct {
	repo adminrepo.UmsAdminRepo
	conf *viper.Viper
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
		Passwd:    hash.Gen(reqData.PassWord),
		RegIpAddr: reqData.RegIpAddr,
	}

	return us.repo.Save(c, admin)
}

// UmsAdminLogin implements UmsAdminService.
func (us *umsAdminService) UmsAdminLogin(c context.Context, reqData *model.UmsAdminLoginReq) (*model.UmsAdminLoginResp, error) {
	admin, exists, err := us.repo.SelectByUserName(c, reqData.UserName)
	if !exists {
		return nil, e.ErrNotFound().WithMsg("用户不存在, 可能未注册")
	}
	if err != nil {
		return nil, err
	}

	if !hash.Check(reqData.PassWord, admin.Passwd) {
		return nil, e.ErrBadRequest().WithMsg("账号或密码错误")
	}

	// 生成Token
	claims := jwt.CustomClaims{
		UserId:   admin.ID,
		UserName: admin.UserName,
	}
	expiresAt := time.Now().Add(time.Minute * 10)
	secret := us.conf.GetString("jwt.admin.secret")
	token, _ := jwt.GenToken(claims, expiresAt, secret)
	respData := &model.UmsAdminLoginResp{
		ID:       admin.ID,
		UserName: admin.UserName,
		NickName: admin.NickName,
		Token:    token,
	}

	return respData, nil
}

func NewUmsAdminService(repo adminrepo.UmsAdminRepo, conf *viper.Viper) UmsAdminService {
	return &umsAdminService{repo, conf}
}
