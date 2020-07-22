package models

type Admin struct {
	Model
	Username string `gorm:"type:VARCHAR(20);NOT NULL" json:"username"`
	Nickname string `gorm:"type:VARCHAR(50);NOT NULL" json:"nickname"`
	Password string `gorm:"type:VARCHAR(32);NOT NULL" json:"password"`
	Salt     string `gorm:"type:VARCHAR(30);NOT NULL" json:"salt"`
	Avatar   string `gorm:"type:VARCHAR(255);NOT NULL" json:"avatar"`
	Email    string `gorm:"type:VARCHAR(100);NOT NULL" json:"email"`
	Token    string `gorm:"type:VARCHAR(59);NOT NULL" json:"token"`
	Status   string `gorm:"type:VARCHAR(30);NOT NULL" json:"status"`
}

func (t *Admin) TableName() string {
	return "gg_admin"
}
