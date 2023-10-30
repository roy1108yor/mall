package model

import "time"

type UmsAdmin struct {
	ID          uint      `xorm:"not null pk autoincr BIGINT(20) id"`
	CreatedAt   time.Time `xorm:"created TIMESTAMP created_at"`
	UpdatedAt   time.Time `xorm:"updated TIMESTAMP updated_at"`
	DeletedAt   time.Time `xorm:"deleted DATETIME deleted_at"`
	UserName    string    `xorm:"not null unique VARCHAR(50) user_name comment('管理员登录账号')"`
	NickName    string    `xorm:"null VARCHAR(50) nick_name comment('管理员昵称')"`
	Passwd      string    `xorm:"not null VARCHAR(60) passwd comment('管理员登录密码')"`
	RegIpAddr   string    `xorm:"null VARCHAR(15) reg_ip_addr comment('注册时候的IP地址')"`
	LoginIpAddr string    `xorm:"null VARCHAR(15) login_ip_addr comment('最后一次登录的IP地址')"`
	LoginTime   time.Time `xorm:"null DATETIME login_time comment('最后一次登录时间')"`
}

type UmsAdminLoginReq struct {
	UserName    string    `json:"user_name" validate:"required" message:"required:{field} 必填"`
	PassWord    string    `json:"password" validate:"required" message:"required:{field} 必填"`
	LoginIpAddr string    `json:"-"`
	LoginTime   time.Time `json:"-"`
}

type UmsAdminRegisterReq struct {
	UserName  string `json:"user_name" validate:"required|min_len:5" message:"required:{field} 必填|min_len:{field} 不能少于5个字符"`
	PassWord  string `json:"password" validate:"required|min_len:6|max_len:20" message:"required:{field} 必填|min_len:{field} 不能少于6个字符|max_len:{field} 不能超过20个字符"`
	RegIpAddr string `json:"-"`
}

type UmsAdminLoginResp struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Token    string `json:"token"`
}

type AllocRoleForAdminReq struct {
	RoleIds []uint `json:"roleIds"`
	UserId  uint   `json:"userId"`
}

func (u *UmsAdmin) TableName() string {
	return "t_ums_admin"
}
