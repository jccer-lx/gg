package models

type SystemLog struct {
	Model
	AdminId  uint   `gorm:"column:admin_id;default:0;NOT NULL" json:"admin_id"` // 管理员id
	Path     string `gorm:"column:path;NOT NULL" json:"path"`                   // 链接
	Page     string `gorm:"column:page;NOT NULL" json:"page"`                   // 行为
	Method   string `gorm:"column:method;NOT NULL" json:"method"`               // 访问类型
	Params   string `gorm:"column:params;NOT NULL" json:"params"`               // 绑定参数内容
	Response string `gorm:"column:response;NOT NULL" json:"params"`             // 返回值
	Ip       string `gorm:"column:ip;NOT NULL" json:"ip"`                       // 登录IP
	FlowIn   int64  `gorm:"column:flow_in;NOT NULL" json:"flow_in"`             // 流量（入）bytes
	FlowOut  int64  `gorm:"column:flow_out;NOT NULL" json:"flow_out"`           // 流量（出）bytes
}

func (m *SystemLog) TableName() string {
	return "gg_system_log"
}
