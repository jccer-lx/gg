package impl

type DataField struct {
	Name      string `json:"name"`
	Title     string `json:"title"`
	JsonTitle string `json:"json_title"`
}

type AdminModelImpl interface {
	TableName() string
	//列表查询时使用的集合
	GormFindOut() interface{}
	//结构表格展示字段
	GetTableFields() ([]*DataField, error)
}
