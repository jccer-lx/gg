package params

import "encoding/json"

type GoodsAddParams struct {
	MainImage    string      `json:"main_image" validate:"required" label:"主图"`
	SliderImage  string      `json:"slider_image"`
	Name         string      `json:"name" validate:"required,min=1,max=20"`
	MainInfo     string      `json:"main_info"`
	Keyword      string      `json:"keyword"`
	BarCode      string      `json:"bar_code"`
	CategoryId   json.Number `json:"category_id" validate:"required"`
	Price        json.Number `json:"price" validate:"required"`
	VipPrice     json.Number `json:"vip_price"`
	OtPrice      json.Number `json:"ot_price"`
	Postage      json.Number `json:"postage"`
	UnitName     string      `json:"unit_name"`
	Sort         json.Number `json:"sort"`
	Sales        json.Number `json:"sales"`
	Stock        json.Number `json:"stock"`
	IsShow       string      `json:"is_show"`
	IsHot        string      `json:"is_hot"`
	IsBenefit    string      `json:"is_benefit"`
	IsBest       string      `json:"is_best"`
	IsNew        string      `json:"is_new"`
	IsPostage    string      `json:"is_postage"`
	GiveIntegral json.Number `json:"give_integral"`
	Cost         json.Number `json:"cost"`
	IsGood       string      `json:"is_good"`
	VirtualSales json.Number `json:"virtual_sales"`
	Browse       json.Number `json:"browse"`
}

func (p *GoodsAddParams) NewParams() GGParams {
	return new(GoodsAddParams)
}

type UpdateGoodsForFieldParams struct {
	ID    uint        `json:"id" validate:"required"`
	Field string      `json:"field" validate:"required"`
	Data  interface{} `json:"data"`
}

func (p *UpdateGoodsForFieldParams) NewParams() GGParams {
	return new(UpdateGoodsForFieldParams)
}
