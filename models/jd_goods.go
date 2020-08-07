package models

type JdGoods struct {
	SkuID    string `gorm:"not null;unique" json:"sku_id"`
	ItemDate string `gorm:"not null;unique" json:"item_date"`
	//品牌
	BrandID     string `json:"brand_id"`
	BrandName   string `json:"brand_name"`
	PName       string `json:"p_name"`
	SkuName     string `json:"sku_name"`
	ProductArea string `json:"product_area"`
	SaleUnit    string `json:"sale_unit"`
	//分类id集合
	CategoryJson string `json:"category_json"`
	//包装尺寸(mm)
	Length int64 `json:"length"`
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
	//规格
	SalePropJson    string `json:"sale_prop_json"`
	SalePropSeqJson string `json:"sale_prop_seq_json"`
	//新版规格
	NewColorSizeJson string `json:"new_color_size_json"`
	//jd价格
	//"l": "",    //划线价
	//"m": "10998.00",
	//"nup": "",
	//"op": "5149.00",
	//"p": "4999.00",
	PriceL  string `json:"price_l"`
	PriceM  string `json:"price_m"`
	PriceOP string `json:"price_o_p"`
	PriceP  string `json:"price_p"`
	//库存情况
	StockState     int64  `json:"stock_state"`
	StockStateName string `json:"stock_state_name"`
	StockAreaJson  string `json:"stock_area_json"`
}

func (m *JdGoods) TableName() string {
	return "gg_jd_item"
}
