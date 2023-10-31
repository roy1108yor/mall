package model

import (
	"time"
)

type PmsProduct struct {
	ID                    uint   `xorm:"not null pk autoincr BIGINT(20) id"`
	BrandId               uint   `xorm:"not null BIGINT(29) brand_id"`
	ProductCategoryId     uint   `xorm:"not null BIGINT(20) product_category_id"`
	ProductAttrCategoryId uint   `xorm:"not null BIGINT(20) product_attr_category_id"`
	Name                  string `xorm:"not null VARCHAR(200) name"`
	Pic                   string `xorm:"not null default null VARCHAR(255) pic"`
	ProductSn             string `xorm:"not null default null VARCHAR(64) product_sn"`
	PublishStatus         int
	IsNew                 int
	IsRecommand           int
	IsVerify              int
	Sort                  int
	OriginPrice           int
	PromotionPrice        int
	GiftGrowth            int
	GiftPoint             int
	SubTitle              string
	Description           string
	Stock                 int
	LowStock              int
	Uint                  int
	Weight                int
	CreatedAt             time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt             time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt             time.Time `xorm:"deleted DATETIME deleted_at"`
}

func (r *PmsProduct) TableName() string {
	return "t_pms_product"
}
