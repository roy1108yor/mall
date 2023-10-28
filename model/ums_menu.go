package model

import "time"

type UmsMenu struct {
	ID        uint      `xorm:"not null pk autoincr BIGINT(20) id comment('唯一标识')"`
	ParentID  uint      `xorm:"no null BIGINT(20) parent_id comment('父级菜单ID')"`
	Name      string    `xorm:"not null VARCHAR(50) name comment('菜单名称')"`
	Level     int       `xorm:"not null INT level comment('菜单等级')"`
	Sort      int       `xorm:"not null TINYINT(1) sort comment('排序')"`
	Icon      string    `xorm:"not null VARCHAR(10) icon comment('菜单图标')"`
	Hidden    int       `xorm:"not null TINYINT(1) hidden comment('是否显示')"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at comment('创建时间')"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at comment('更新时间')"`
	DeletedAt time.Time `xorm:"deleted DATETIME deleted_at comment('删除时间')"`
}

type AddUmsMenuReq struct {
	ParentId uint   `json:"parentId" validate:"required|int" message:"required:{field} 必填|int:{field} 必须是数字类型"`
	Name     string `json:"name" validate:"required" message:"required:{field} 必填"`
	Icon     string `json:"icon" validate:"required" message:"required:{field} 必填"`
	Hidden   int    `json:"hidden" validate:"required" message:"required:{field} 必填"`
	Sort     int    `json:"sort" validate:"required|int" message:"required:{field} 必填|int:{field} 必须是数字类型"`
}

func (r *UmsMenu) TableName() string {
	return "t_ums_menu"
}

func (r *AddUmsMenuReq) ToModel() *UmsMenu {
	return &UmsMenu{
		ParentID: r.ParentId,
		Name:     r.Name,
		Icon:     r.Icon,
		Hidden:   r.Hidden,
		Sort:     r.Sort,
	}
}