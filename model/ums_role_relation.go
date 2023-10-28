package model

import "time"

type UmsRoleRelation struct {
	ID        uint      `xorm:"not null pk autoincr BIGINT(20) id comment('唯一标识')"`
	AdminID   uint      `xorm:"not null BIGINT(20) admin_id comment('管理员ID')"`
	RoleID    uint      `xorm:"not null BITINT(20) role_id comment('角色ID')"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at comment('创建时间')"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at comment('更新时间')"`
}

func (r *UmsRoleRelation) TableName() string {
	return "t_ums_role_relation"
}

type AddUmsRoleRelationReq struct {
	AdminID uint `json:"adminId" validate:"required|int" message:"required:{field} 必填|int{field} 必须是数字类型"`
	RoleID  uint `json:"roleId" validate:"required|int" message:"required:{field} 必填|int{field} 必须是数字类型"`
}

func (r *AddUmsRoleRelationReq) ToModel() *UmsRoleRelation {
	return &UmsRoleRelation{
		AdminID: r.AdminID,
		RoleID:  r.RoleID,
	}
}
