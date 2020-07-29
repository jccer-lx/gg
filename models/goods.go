package models

type Goods struct {
	Model
	AdminId         uint    `gorm:"column:admin_id;NOT NULL" json:"admin_id"`                                   // 管理员
	Image           string  `gorm:"column:image;NOT NULL" json:"image"`                                         // 商品图片
	SliderImageJson string  `gorm:"column:slider_image_json;NOT NULL" json:"slider_image_json"`                 // 轮播图
	Name            string  `gorm:"column:name;NOT NULL" json:"name"`                                           // 商品名称
	MainInfo        string  `gorm:"column:main_info;NOT NULL" json:"main_info"`                                 // 商品简介
	Keyword         string  `gorm:"column:keyword;NOT NULL" json:"keyword"`                                     // 关键字
	BarCode         string  `gorm:"column:bar_code;NOT NULL" json:"bar_code"`                                   // 商品条码（一维码）
	CategoryId      string  `gorm:"column:category_id;NOT NULL" json:"category_id"`                             // 分类id
	Price           float64 `gorm:"column:price;type:decimal(10,2);default:0.00;NOT NULL" json:"price"`         // 商品价格
	VipPrice        float64 `gorm:"column:vip_price;type:decimal(10,2);default:0.00;NOT NULL" json:"vip_price"` // 会员价格
	OtPrice         float64 `gorm:"column:ot_price;type:decimal(10,2);default:0.00;NOT NULL" json:"ot_price"`   // 市场价
	Postage         float64 `gorm:"column:postage;type:decimal(10,2);default:0.00;NOT NULL" json:"postage"`     // 邮费
	UnitName        string  `gorm:"column:unit_name;NOT NULL" json:"unit_name"`                                 // 单位名
	Sort            int     `gorm:"column:sort;default:0;NOT NULL" json:"sort"`                                 // 排序
	Sales           int     `gorm:"column:sales;default:0;NOT NULL" json:"sales"`                               // 销量
	Stock           int     `gorm:"column:stock;default:0;NOT NULL" json:"stock"`                               // 库存
	IsShow          int     `gorm:"column:is_show;default:1;NOT NULL" json:"is_show"`                           // 状态（0：未上架，1：上架）
	IsHot           int     `gorm:"column:is_hot;default:0;NOT NULL" json:"is_hot"`                             // 是否热卖
	IsBenefit       int     `gorm:"column:is_benefit;default:0;NOT NULL" json:"is_benefit"`                     // 是否优惠
	IsBest          int     `gorm:"column:is_best;default:0;NOT NULL" json:"is_best"`                           // 是否精品
	IsNew           int     `gorm:"column:is_new;default:0;NOT NULL" json:"is_new"`                             // 是否新品
	IsPostage       int     `gorm:"column:is_postage;default:0;NOT NULL" json:"is_postage"`                     // 是否包邮
	GiveIntegral    string  `gorm:"column:give_integral;NOT NULL" json:"give_integral"`                         // 获得积分
	Cost            float64 `gorm:"column:cost;type:decimal(10,2);NOT NULL" json:"cost"`                        // 成本价
	IsSecKill       int     `gorm:"column:is_sec_kill;default:0;NOT NULL" json:"is_sec_kill"`                   // 秒杀状态 0 未开启 1已开启
	IsBargain       int     `gorm:"column:is_bargain" json:"is_bargain"`                                        // 砍价状态 0未开启 1开启
	IsGood          int     `gorm:"column:is_good;default:0;NOT NULL" json:"is_good"`                           // 是否优品推荐
	VirtualSales    int     `gorm:"column:virtual_sales;default:100" json:"virtual_sales"`                      // 虚拟销量
	Browse          int     `gorm:"column:browse;default:0" json:"browse"`                                      // 浏览量
	CodePath        string  `gorm:"column:code_path;NOT NULL" json:"code_path"`                                 // 商品二维码地址(用户小程序海报)
	VideoLink       string  `gorm:"column:video_link;NOT NULL" json:"video_link"`                               // 主图视频链接
	TempId          uint    `gorm:"column:temp_id;default:1;NOT NULL" json:"temp_id"`                           // 运费模板ID
}

func (m *Goods) TableName() string {
	return "gg_goods"
}
