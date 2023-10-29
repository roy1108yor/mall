package model

import "time"

type UmsRole struct {
	ID          uint      `xorm:"not null pk autoincr BIGINT(20) id comment('唯一标识')"`
	Name        string    `xorm:"not null VARCHAR(100) name comment('名称')"`
	Description string    `xorm:"null default null VARCHAR(500) description comment('描述')"`
	AdminCount  int       `xorm:"null default null INT(11) admin_count comment('后台用户数量')"`
	Status      int       `xorm:"not null default 1 TINYINT(1) status comment('状态: 1正常 0禁用')"`
	Sort        int       `xorm:"not null default 0 TINYINT(1) sort comment('排序, 数字越小越靠前')"`
	CreatedAt   time.Time `xorm:"created TIMESTAMP created_at comment('创建时间')"`
	UpdatedAt   time.Time `xorm:"updated TIMESTAMP updated_at comment('更新时间')"`
	DeletedAt   time.Time `xorm:"deleted DATETIME deleted_at comment('删除时间')"`
}

type AddUmsRoleReq struct {
	Name        string `json:"name" validate:"required|min_len:2" message:"required:{field} 必填|min_len:{field} 不能少于2个字符"`
	Description string `json:"description"`
	Status      int    `json:"status" validate:"required|int|min:0|max:1" message:"required:{field} 必填|int:{field} 必须是整数类型|min:{field} 应该是0或1|max:{field} 应该是0或1"`
}

func (r *UmsRole) TableName() string {
	return "t_ums_role"
}

func (r *AddUmsRoleReq) ToModel() *UmsRole {
	return &UmsRole{
		Name:        r.Name,
		Description: r.Description,
		Status:      r.Status,
	}
}
