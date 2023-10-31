package model

import "time"

type UmsRoleRelation struct {
	ID        uint      `xorm:"not null pk autoincr BIGINT(20) id comment('唯一标识')"`
	AdminId   uint      `xorm:"not null BIGINT(20) admin_id comment('管理员ID')"`
	RoleId    uint      `xorm:"not null BITINT(20) role_id comment('角色ID')"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at comment('创建时间')"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at comment('更新时间')"`
	DeletedAt time.Time `xorm:"deleted DATETIME deleted_at comment('删除时间')"`
}

func (r *UmsRoleRelation) TableName() string {
	return "t_ums_role_relation"
}

type UmsRoleRelationInReq struct {
	AdminId uint   `json:"adminId" validate:"uint" message:"uint{field} 必填且是数字类型"`
	RoleIds []uint `json:"roleIds"`
}
