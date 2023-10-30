package model

import "time"

type UmsResource struct {
	ID          uint      `xorm:"not null pk autoincr BIGINT(20) id"`
	Name        string    `xorm:"not null VARCHAR(200) name comment('资源名称')"`
	URL         string    `xorm:"not null VARCHAR(500) url comment('资源路径')"`
	Description string    `xorm:"not null VARCHAR(500) description comment('资源描述')"`
	CategoryID  uint      `xorm:"not null BITINT(20) category_id comment('资源分类ID')"`
	CreatedAt   time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt   time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt   time.Time `xorm:"deleted DATETIME deleted_at"`
}

type UmsResourceInReq struct {
	Name        string `json:"name" validate:"required" message:"required:{field} 必填项"`
	URL         string `json:"url" validate:"required" message:"required:{field} 必填项"`
	Description string `json:"description" validate:"required" message:"required:{field} 必填项"`
	CategoryId  uint   `json:"categoryId" validate:"uint|min:0" message:"uint:{field} 必填且必须是数字类型|min:{field} 不能小于0"`
}

func (r *UmsResource) TableName() string {
	return "t_ums_resource"
}

func (r *UmsResourceInReq) ToModel() *UmsResource {
	return &UmsResource{
		Name:        r.Name,
		URL:         r.URL,
		Description: r.Description,
		CategoryID:  r.CategoryId,
	}
}
