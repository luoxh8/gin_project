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
		for i := 0; i < 5; i++ {
			o.Insert(&models.GoodsBannerPics{
				GoodsId: value.GoodsId,
				Url:     "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1540726838364&di=5e7b11914809a739bba934cfd42c2b47&imgtype=0&src=http%3A%2F%2Fimgsrc.baidu.com%2Fimgad%2Fpic%2Fitem%2Faa18972bd40735fa7b41418395510fb30f240852.jpg",
				Index:   i,
			})
		}
	}

}
