package model

import "time"

type UmsRoleMenuRelation struct {
	ID        uint      `xorm:"not null pk autoincr BIGINT(20) id comment('唯一标识')"`
	RoleId    uint      `xorm:"not null BIGINT(20) role_id"`
	MenuId    uint      `xorm:"not null BITINT(20) menu_id"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at comment('创建时间')"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at comment('更新时间')"`
	DeletedAt time.Time `xorm:"deleted DATETIME deleted_at comment('删除时间')"`
}

type UmsRoleMenuRelationInReq struct {
	RoleId uint `validate:"required|uint" json:"roleId"`
	MenuId uint `validate:"required|uint" json:"menuId"`
}

func (r *UmsRoleMenuRelationInReq) ToModel() *UmsRoleMenuRelationInReq {
	return &UmsRoleMenuRelationInReq{
		RoleId: r.RoleId,
		MenuId: r.MenuId,
	}
}

func (r *UmsRoleMenuRelation) TableName() string {
	return "t_ums_role_menu_relation"
}
