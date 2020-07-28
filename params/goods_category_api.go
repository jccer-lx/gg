package params

import "encoding/json"

type GoodsCategoryParams struct {
	Pid      json.Number `gorm:"column:pid;NOT NULL" json:"pid"`             // 父id
	CateName string      `json:"cate_name" validate:"required,min=1,max=10"` // 分类名称
	Sort     json.Number `json:"sort"`                                       // 排序
	Pic      string      `json:"pic"`                                        // 图标
	IsShow   string      `json:"is_show"`                                    // 是否推荐
}

func (p *GoodsCategoryParams) NewParams() GGParams {
	return new(GoodsCategoryParams)
}

type GoodsCategoryUpdateParams struct {
	CateName string      `json:"cate_name" validate:"required,min=1,max=10"` // 分类名称
	Sort     json.Number `json:"sort"`                                       // 排序
	Pic      string      `json:"pic"`                                        // 图标
	IsShow   json.Number `json:"is_show"`                                    // 是否推荐
}

func (p *GoodsCategoryUpdateParams) NewParams() GGParams {
	return new(GoodsCategoryUpdateParams)
}
