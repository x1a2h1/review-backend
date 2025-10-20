package models

const (
	UserStatusNormal int8 = iota
	UserStatusDisable
)

func ExistsUserStatus(status int8) bool {
	if status == UserStatusNormal || status == UserStatusDisable {
		return true
	}
	return false
}

type User struct {
	Base
	Account   string `json:"account" gorm:"type:varchar(50);index:idx_account;unique;comment:账号"`
	Nickname  string `json:"nickname" gorm:"type:varchar(50);index:idx_nickname;comment:昵称、别名"`
	Avatar    string `json:"avatar" gorm:"type:varchar(255);comment:头像"`
	Password  string `json:"-" gorm:"type:varchar(255);not null;comment:用户密码"`
	Cellphone string `json:"cellphone" gorm:"type:varchar(50);index:idx_cellphone;comment:手机号"`
	Email     string `json:"email" gorm:"type:varchar(100);index:idx_email;comment:邮箱"`
	Status    int8   `json:"status" gorm:"type:tinyint(4);default:0;comment:账号状态，0-正常，1-禁用"`
	Root      bool   `json:"root" gorm:"<-:create;type:tinyint(4);default:0;comment:是否超级用户，0-否，1-是"`
	Role      string `json:"role" gorm:"type:varchar(50);comment:用户角色"`
}
