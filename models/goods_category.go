package models

const (
	GoodsCategoryIsShowTrue  = 1
	GoodsCategoryIsShowFalse = 0
)

type GoodsCategory struct {
	Model
	Pid       uint             `gorm:"column:pid;NOT NULL" json:"pid"`                   // 父id
	PCateName string           `gorm:"-" json:"p_cate_name"`                             // 父cate_name
	CateName  string           `gorm:"column:cate_name;NOT NULL" json:"cate_name"`       // 分类名称
	Sort      int              `gorm:"column:sort;NOT NULL" json:"sort"`                 // 排序
	Pic       string           `gorm:"column:pic;NOT NULL" json:"pic"`                   // 图标
	IsShow    int              `gorm:"column:is_show;default:1;NOT NULL" json:"is_show"` // 是否推荐
	Children  []*GoodsCategory `gorm:"-" json:"children"`                                //子菜单
}

func (m *GoodsCategory) TableName() string {
	return "gg_goods_category"
}
