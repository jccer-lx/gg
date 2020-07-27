package models

const (
	AuthRuleTypeMenu     = `menu`
	AuthRuleTypeFile     = `file`
	AuthRuleIsMenuTrue   = 1
	AuthRuleIsMenuFalse  = 0
	AuthRuleStatusNormal = "normal"
)

type AuthRule struct {
	Model
	Pid    uint   `gorm:"column:pid;NOT NULL" json:"pid"`                // 父ID
	PTitle string `gorm:"-" json:"p_title"`                              // 父title
	Type   string `gorm:"column:type;NOT NULL" json:"type"`              // 类型
	Name   string `gorm:"column:name;NOT NULL;unique_index" json:"name"` // 规则值（路由）
	Title  string `gorm:"column:title;NOT NULL" json:"title"`            // 规则名称
	Remark string `gorm:"column:remark;NOT NULL" json:"remark"`          // 备注
	IsMenu int    `gorm:"column:is_menu;NOT NULL" json:"is_menu"`        // 是否为菜单
	Weigh  int    `gorm:"column:weigh;NOT NULL" json:"weigh"`            // 权重
	Status string `gorm:"column:status;NOT NULL" json:"status"`          // 状态
}

func (t *AuthRule) TableName() string {
	return "gg_auth_rule"
}
