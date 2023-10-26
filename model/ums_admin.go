package model

import "time"

type UmsAdmin struct {
	ID        uint      `xorm:"not null pk autoincr BIGINT(20) id"`
	UserName  string    `xorm:"not null unique VARCHAR(50) user_name"`
	NickName  string    `xorm:"null VARCHAR(50) nick_name"`
	Passwd    string    `xorm:"not null VARCHAR(60) passwd"`
	Email     string    `xorm:"not null unique VARCHAR(100) email"`
	RegIpAddr string    `xorm:"null VARCHAR(15) reg_ip_addr"`
	LoginTime time.Time `xorm:"null DATETIME login_time"`
	CreatedAt time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt time.Time `xorm:"deleted TIMESTAMP deleted_at"`
}

func (UmsAdmin) TableName() string {
	return "t_ums_admin"
}
