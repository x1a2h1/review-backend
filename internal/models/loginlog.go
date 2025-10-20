package models

type LoginLog struct {
	Base
	UserId    uint   `json:"userId" gorm:"type:bigint(20);comment:登录用户的ID"`
	Account   string `json:"account" gorm:"type:varchar(50);comment:账户"`
	Nickname  string `json:"nickname" gorm:"type:varchar(50);not null;index:idx_uname;comment:昵称"`
	Ip        string `json:"ip" gorm:"type:varchar(50);comment:登录IP"`
	Address   string `json:"address" gorm:"type:varchar(255);comment:登录地点"`
	UserAgent string `json:"userAgent" gorm:"type:varchar(255);comment:浏览器的userAgent"`
}

func (*LoginLog) TableName() string {
	return "login_log"
}
