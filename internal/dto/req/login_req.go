package req

type LoginReq struct {
	Account  string `json:"account" binding:"required" label:"账号"`
	Password string `json:"password" binding:"required,max=30" label:"密码"`
}
