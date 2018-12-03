package main

import (
	`gin_project/core/h_categories`
	`gin_project/models`
	`github.com/astaxie/beego/orm`
)

func main() {
	var (
		goods []*models.Goods
	)

	orm.NewOrm().QueryTable("goods").All(&goods)
	for _, value := range goods {
		orm.NewOrm().Insert(&models.GoodsRoomCount{
			GoodsId:   value.GoodsId,
			Onlookers: int64(h_categories.GenRandomInt(1000, 2000)),
			Bidders:   int64(h_categories.GenRandomInt(500, 1000)),
		})
	}
}
