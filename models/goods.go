package models

import (
	`github.com/astaxie/beego/orm`
)

type PayMan struct {
	Base
	Uid        string  `orm:"size(7);index;unique" json:"uid"`
	Nick       string  `orm:"size(20)" json:"nick"`                    // 出价的人
	IsOut      bool    `orm:"default(true)" json:"is_out"`             // 是否出局
	Area       string  `orm:"size(10)" json:"area"`                    // 地区
	PayPoint   float64 `orm:"digits(12);decimals(2)" json:"pay_point"` // 当时出价的金币
	TimeRemain int     `json:"time_remain"`                            // 剩余时间
}

type Category struct {
	Base
	Name        string `orm:"size(30);" json:"name"`
	Description string `orm:"size(30);type(text)" json:"description"`
	IsTab       bool   `orm:"default(false)" json:"is_tab"`
}

type Goods struct {
	Base
	GoodsId         string  `orm:"size(64);index;unique" json:"goods_id"`       // 商品的id
	OwnerId         string  `orm:"size(7)" json:"owner_id"`                     // 0 代表没有Owner
	CategoryId      int     `json:"category_id"`                                // 分类id
	Avatar          string  `orm:"size(1024)" json:"avatar"`                    // 商品缩略图
	TotalPrice      float64 `orm:"digits(12);decimals(2)" json:"total_price"`   // 商品总价
	CurrentPrice    float64 `orm:"digits(12);decimals(2)" json:"current_price"` // 当前出价
	Name            string  `orm:"size(50)" json:"name"`                        // 商品名称
	PayRecordsCount int     `json:"pay_records_count"`                          // 出价次数
	CostMoney       int     `json:"cost_money"`                                 // 消耗金币
	Period          int     `json:"period"`                                     // 期数
	Like            bool    `orm:"default(true)" json:"like"`                   // 类型
	IsEnd           bool    `orm:"default(true)" json:"is_end"`                 // 是否结束

	/**
		dict字段，无实际作用
	*/
	GoodsBannerPics interface{} `orm:"-" json:"goods_banner_pics"`
	Onlookers       interface{} `orm:"-" json:"onlookers"`
	Bidders         interface{} `orm:"-" json:"bidders"`
}

func (g *Goods) SaveCurrentPrice() bool {
	if _, err := orm.NewOrm().Update(g.CurrentPrice); err != nil {
		return false
	}
	return true
}

func (g *Goods) FindByGoodsId(goodsId string) bool {
	if err := orm.NewOrm().QueryTable("goods").Filter("goods_id", goodsId).One(g); err != nil {
		return false
	}
	return true
}

type GoodsBannerPics struct {
	Base
	GoodsId string `orm:"size(64)" json:"-"`
	Index   int    `json:"index"`
	Url     string `orm:"size(1024)" json:"url"`
}

type GoodsPayMans struct {
	Base
	GoodsId  string `orm:"size(64);index;unique" json:"goods_id"`  // 商品的id
	PayManId string `orm:"size(7);index;unique" json:"pay_man_id"` // 用户id
}

type RealTimeAuction struct {
	Base
	GoodsId   string `orm:"size(64);index;unique" json:"goods_id"`  // 商品的id
	PayManId  string `orm:"size(7);index;unique" json:"pay_man_id"` // 用户id
	TimeLimit int    `json:"time_limit"`
}

type GoodsRoomCount struct {
	Base
	GoodsId   string `orm:"size(64)" json:"-"`
	Onlookers int64  `json:"onlookers"`
	Bidders   int64  `json:"bidders"`
}

func (grc *GoodsRoomCount) Count(goodsId string) bool {
	if err := orm.NewOrm().QueryTable("goods_room_count").Filter("goods_id", goodsId).One(&grc); err != nil {
		return false
	}
	return true
}
