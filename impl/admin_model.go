package impl

type DataField struct {
	Name      string `json:"name"`
	Title     string `json:"title"`      //标题
	JsonTitle string `json:"json_title"` //html name/id 使用
	Type      string //表单类型
}

type AdminModelImpl interface {
	TableName() string
	//列表查询时使用的集合
	GormFindOut() interface{}
	//结构表格展示字段
	GetTableFields() ([]*DataField, error)
}
