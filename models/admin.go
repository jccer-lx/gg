package models

const AdminStatusNormal = "normal"

type Admin struct {
	Model
	Username string `gorm:"type:varchar(20);NOT NULL" json:"username"`
	Nickname string `gorm:"type:varchar(50);NOT NULL" json:"nickname"`
	Password string `gorm:"type:varchar(32);NOT NULL" json:"password"`
	Salt     string `gorm:"type:varchar(30);NOT NULL" json:"salt"`
	Avatar   string `gorm:"type:varchar(255);NOT NULL" json:"avatar"`
	Email    string `gorm:"type:varchar(100);NOT NULL" json:"email"`
	Token    string `gorm:"type:varchar(59);NOT NULL" json:"token"`
	Status   string `gorm:"type:varchar(30);NOT NULL" json:"status"`
}

func (t *Admin) TableName() string {
	return "gg_admin"
}
