package model

import (
	"time"
)

type PmsProduct struct {
	ID              uint   `xorm:"not null pk autoincr BIGINT(20) id"`
	BrandId         uint   `json:"brand_id" xorm:"not null BIGINT(29) brand_id"`
	CategoryId      uint   `json:"category_id" xorm:"not null BIGINT(20) category_id"`
	AttrCategoryId  uint   `json:"attr_category_id" xorm:"not null BIGINT(20) attr_category_id"`
	Name            string `json:"name" xorm:"not null VARCHAR(200) name"`
	Description     string `json:"description" xorm:"null default null TEXT"`
	SubTitle        string `json:"subTitle" xorm:"null default null VARCHAR(255) sub_title"`
	Pic             string `json:"pic" xorm:"null default null VARCHAR(255) pic"`
	ProductSno      string `json:"product_sno" xorm:"null default null VARCHAR(64) product_sn0"`
	PublishStatus   int    `json:"publish_status" xorm:"null default null TINYING(1) publish_status"`
	NewStatus       int    `json:"new_status" xorm:""`
	RecommendStatus int
	VerifyStatus    int
	Sort            int
	OriginPrice     int
	PromotionPrice  int
	GiftGrowth      int
	GiftPoint       int
	Stock           int
	LowStock        int
	Uint            int
	Weight          int
	CreatedAt       time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt       time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt       time.Time `xorm:"deleted DATETIME deleted_at"`
}

func (r *PmsProduct) TableName() string {
	return "t_pms_product"
}
