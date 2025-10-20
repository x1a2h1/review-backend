package req

type UserUpsertReq struct {
	Account   string `json:"account" binding:"required,min=3,max=50" label:"账号"`
	Nickname  string `json:"nickname" binding:"omitempty,min=1,max=50" label:"昵称"`
	Avatar    string `json:"avatar" binding:"omitempty,url,max=255" label:"头像"`
	Email     string `json:"email" binding:"omitempty,email,max=100" label:"邮箱"`
	Cellphone string `json:"cellphone" binding:"omitempty,max=11" label:"手机号"`
}
