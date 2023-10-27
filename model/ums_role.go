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

func (UmsRole) TableName() string {
	return "t_ums_role"
}
