package models

import (
	"time"
)

//fa的会员表
type User struct {
	Id             int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                 // ID
	GroupId        int       `gorm:"column:group_id;default:0;NOT NULL" json:"group_id"`             // 组别ID
	Username       string    `gorm:"column:username;NOT NULL" json:"username"`                       // 用户名
	Nickname       string    `gorm:"column:nickname;NOT NULL" json:"nickname"`                       // 昵称
	Password       string    `gorm:"column:password;NOT NULL" json:"password"`                       // 密码
	Salt           string    `gorm:"column:salt;NOT NULL" json:"salt"`                               // 密码盐
	Email          string    `gorm:"column:email;NOT NULL" json:"email"`                             // 电子邮箱
	Mobile         string    `gorm:"column:mobile;NOT NULL" json:"mobile"`                           // 手机号
	Avatar         string    `gorm:"column:avatar;NOT NULL" json:"avatar"`                           // 头像
	Level          int       `gorm:"column:level;default:0;NOT NULL" json:"level"`                   // 等级
	Gender         int       `gorm:"column:gender;default:0;NOT NULL" json:"gender"`                 // 性别
	Birthday       time.Time `gorm:"column:birthday" json:"birthday"`                                // 生日
	Bio            string    `gorm:"column:bio;NOT NULL" json:"bio"`                                 // 格言
	Money          string    `gorm:"column:money;default:0.00;NOT NULL" json:"money"`                // 余额
	Score          int       `gorm:"column:score;default:0;NOT NULL" json:"score"`                   // 积分
	Successions    int       `gorm:"column:successions;default:1;NOT NULL" json:"successions"`       // 连续登录天数
	Maxsuccessions int       `gorm:"column:maxsuccessions;default:1;NOT NULL" json:"maxsuccessions"` // 最大连续登录天数
	Prevtime       int       `gorm:"column:prevtime" json:"prevtime"`                                // 上次登录时间
	Logintime      int       `gorm:"column:logintime" json:"logintime"`                              // 登录时间
	Loginip        string    `gorm:"column:loginip;NOT NULL" json:"loginip"`                         // 登录IP
	Loginfailure   int       `gorm:"column:loginfailure;default:0;NOT NULL" json:"loginfailure"`     // 失败次数
	Joinip         string    `gorm:"column:joinip;NOT NULL" json:"joinip"`                           // 加入IP
	Jointime       int       `gorm:"column:jointime" json:"jointime"`                                // 加入时间
	Createtime     int       `gorm:"column:createtime" json:"createtime"`                            // 创建时间
	Updatetime     int       `gorm:"column:updatetime" json:"updatetime"`                            // 更新时间
	Token          string    `gorm:"column:token;NOT NULL" json:"token"`                             // Token
	Status         string    `gorm:"column:status;NOT NULL" json:"status"`                           // 状态
	Verification   string    `gorm:"column:verification;NOT NULL" json:"verification"`               // 验证
	Openid         string    `gorm:"column:openid;not null;unique"`
	Unionid        string    `gorm:"column:unionid;"`
}

func (u *User) TableName() string {
	return "zx_user"
}
