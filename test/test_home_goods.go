package main

import (
	`gin_project/models`
	`github.com/astaxie/beego/orm`
)

func main() {
	o := orm.NewOrm()
	var (
		goods []*models.Goods
	)

	o.QueryTable("goods").All(&goods)
	for _, value := range goods {
		o.Insert(&models.HomeGoods{
			GoodsId:      value.GoodsId,
			Title:        value.Name,
			Pic:          value.Avatar,
			CurrentPrice: value.CurrentPrice,
			Like:         true,
		})
	}
}
