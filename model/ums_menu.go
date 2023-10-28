package model

import "time"

type UmsMenu struct {
	ID        uint `xorm:"not null pk autoincr BIGINT(20) id comment('唯一标识')"`
	ParentID  uint
	Title     string
	Name      string
	Level     int
	Sort      int
	Icon      string
	Hidden    int
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at comment('创建时间')"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at comment('更新时间')"`
	DeletedAt time.Time `xorm:"deleted DATETIME deleted_at comment('删除时间')"`
}

func (UmsMenu) TableName() string {
	return "t_ums_menu"
}
