package params

//使用自定义结构体接受参数，防止直接操作model的意外惊喜
type AdminAddApiParams struct {
	AdminUpdateApiParams
	LoginApiParams
	Salt   string `from:"salt" json:"salt"`
	Avatar string `from:"avatar" json:"avatar"`
}

func (p *AdminAddApiParams) NewParams() GGParams {
	return &AdminAddApiParams{}
}

type AdminUpdateApiParams struct {
	Nickname string `from:"nickname" json:"nickname" binding:"required" label:"昵称"`
	Email    string `from:"email" json:"email" validate:"email" label:"邮箱"`
}

func (p *AdminUpdateApiParams) NewParams() GGParams {
	return &AdminUpdateApiParams{}
}

type LoginApiParams struct {
	Username string `from:"username" json:"username" validate:"required,min=6,max=30" label:"账号"`
	Password string `from:"password" json:"password" validate:"required,min=6,max=20" label:"密码"`
}

func (p *LoginApiParams) NewParams() GGParams {
	return &LoginApiParams{}
}
