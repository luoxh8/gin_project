package models

type Banner struct {
	Base
	Index       int    `json:"index"` // 控制轮播循序
	Activity    string `orm:"size(15)" json:"activity"`
	IosActivity string `orm:"size(15)" json:"ios_activity"`
	Url         string `orm:"size(1024)" json:"url"`
	Title       string `orm:"size(126)" json:"title"`
	Pic         string `orm:"size(1024)" json:"pic"`
	Params      string `orm:"size(50)" json:"params"`
}

type SmallIcon struct {
	Base
	Index       int    `json:"index"` // 控制轮播循序
	Activity    string `orm:"size(15)" json:"activity"`
	IosActivity string `orm:"size(15)" json:"ios_activity"`
	Url         string `orm:"size(1024)" json:"url"`
	Title       string `orm:"size(15)" json:"title"`
	Pic         string `orm:"size(1024)" json:"pic"`
	Params      string `orm:"size(50)" json:"params"`
}

type HomeGoods struct {
	Base
	GoodsId      string  `orm:"size(64);index;unique" json:"goods_id"` // 商品的id
	Title        string  `orm:"size(126)" json:"title"`
	Pic          string  `orm:"size(1024)" json:"pic"`
	CurrentPrice float64 `orm:"digits(12);decimals(2)" json:"current_price"`
	Like         bool    `orm:"default(true)" json:"like"`
}
