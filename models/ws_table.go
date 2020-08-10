package models

type WsTable struct {
	Model
	Title string `gorm:"column:title;NOT NULL" json:"title"`
}

func (m *WsTable) TableName() string {
	return "gg_ws_table"
}

type WsTableData struct {
	Model
	F0 string `json:"f0"`
	F1 string `json:"f1"`
	F2 string `json:"f2"`
	F3 string `json:"f3"`
	F4 string `json:"f4"`
	F5 string `json:"f5"`
	F6 string `json:"f6"`
	F7 string `json:"f7"`
	F8 string `json:"f8"`
	F9 string `json:"f9"`
}

func (m *WsTableData) TableName() string {
	return "gg_ws_table_data"
}
