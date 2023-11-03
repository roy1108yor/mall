package adminsrv

import (
	"context"
	"fmt"
	"time"

	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/e"
	"github.com/kalougata/mall/pkg/hash"
	"github.com/kalougata/mall/pkg/jwt"
	adminrepo "github.com/kalougata/mall/repo/admin"
	"github.com/spf13/viper"
)

type umsAdminService struct {
	repo     adminrepo.UmsAdminRepo
	roleRepo adminrepo.UmsRoleRepo
	conf     *viper.Viper
}

type UmsAdminService interface {
	Register(c context.Context, reqData *model.UmsAdminInReq) error
	Login(c context.Context, reqData *model.UmsAdminLoginReq) (respData *model.UmsAdminLoginOut, err error)
	AllocRoleForAdmin(c context.Context, data *model.UmsRoleRelationInReq) error
}

// AllocRoleForAdmin 为后台用户分配角色
func (us *umsAdminService) AllocRoleForAdmin(c context.Context, data *model.UmsRoleRelationInReq) error {
	list := make([]*model.UmsRoleRelation, len(data.RoleIds))
	for i, id := range data.RoleIds {
		list[i] = new(model.UmsRoleRelation)
		list[i].AdminId = data.AdminId
		list[i].RoleId = id
	}
	if _, err := us.roleRepo.BatchInsertRoleRelationForAdmin(c, list); err != nil {
		return e.ErrInternalServer().WithMsg("分配角色失败, 请稍后再试~").WithErr(err)
	}

	return nil
}

// Register 后台用户注册
func (us *umsAdminService) Register(c context.Context, reqData *model.UmsAdminInReq) error {
	_, exists, err := us.repo.SelectByUserName(c, reqData.UserName)
	if err != nil {
		return err
	}
	if exists {
		return e.ErrBadRequest().WithMsg(fmt.Sprintf("UserName: %v 用户名已被注册", reqData.UserName))
	}
	if count, err := us.repo.Create(c, reqData.ToModel()); err != nil || count <= 0 {
		return e.ErrInternalServer().WithMsg("注册失败, 请稍后再试~")
	}

	return nil
}

// Login 后台用户登录
func (us *umsAdminService) Login(c context.Context, reqData *model.UmsAdminLoginReq) (*model.UmsAdminLoginOut, error) {
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
	// secret := us.conf.GetString("jwt.admin.secret")
	token, err := jwt.GenToken(claims, time.Now().Add(time.Minute*10), "admin")
	if err != nil {
		return nil, e.ErrInternalServer().WithMsg("生成token失败")
	}
	respData := &model.UmsAdminLoginOut{
		ID:       admin.ID,
		UserName: admin.UserName,
		NickName: admin.NickName,
		Token:    token,
	}

	return respData, nil
}

func NewUmsAdminService(repo adminrepo.UmsAdminRepo, roleRepo adminrepo.UmsRoleRepo, conf *viper.Viper) UmsAdminService {
	return &umsAdminService{repo, roleRepo, conf}
}
