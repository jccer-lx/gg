package models

type Member struct {
	Model
	Username string  `gorm:"column:username;NOT NULL;" json:"username"`
	Password string  `gorm:"column:password;NOT NULL;" json:"password"`
	Salt     string  `gorm:"column:salt;NOT NULL;" json:"salt"`
	Mobile   string  `gorm:"column:mobile;" json:"mobile"`
	Nickname string  `gorm:"column:nickname;" json:"nickname"`
	Openid   string  `gorm:"column:openid;" json:"openid"`
	Unionid  string  `gorm:"column:unionid;" json:"unionid"`
	Money    float64 `gorm:"column:money;type:decimal(10,2);default:0.00;NOT NULL" json:"money"`
}

func (m *Member) TableName() string {
	return "gg_member"
}
