package model

import "time"

type UmsResourceCategory struct {
	ID        uint      `xorm:"not null pk autoincr BIGINT(20) id"`
	Name      string    `xorm:"not null"`
	Sort      int       `xorm:"not null"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt time.Time `xorm:"deleted DATETIME deleted_at"`
}

type UmsResourceCategoryInReq struct {
	Name string `json:"name" validate:"required" message:"required:{field} 必填项"`
	Sort int    `json:"sort" validate:"int|min:0" message:"int:{field} 必须是数字类型|min:{field} 不能小于0"`
}

func (r *UmsResourceCategory) TableName() string {
	return "t_ums_resource_category"
}

func (r *UmsResourceCategoryInReq) ToModel() *UmsResourceCategory {
	return &UmsResourceCategory{
		Name: r.Name,
		Sort: r.Sort,
	}
}
