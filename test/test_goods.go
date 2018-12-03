package main

import (
	`gin_project/core/h_categories`
	`gin_project/models`
	`github.com/astaxie/beego/orm`
)

func main() {
	o := orm.NewOrm()
	var (
		goods1 models.Goods
		goods2 models.Goods
		goods3 models.Goods
		goods4 models.Goods
		goods5 models.Goods
	)
	goods1.GoodsId = h_categories.GenGoodsId()
	goods1.OwnerId = "0"
	goods1.CategoryId = 0
	goods1.Avatar = "https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=3406060575,3655381718&fm=11&gp=0.jpg"
	goods1.CurrentPrice = 0
	goods1.Like = true
	goods1.CostMoney = 1
	goods1.PayRecordsCount = 0
	goods1.IsEnd = false
	goods1.TotalPrice = 9999
	goods1.Name = "iPhone XS Max 256G"
	goods1.Period = 1

	goods2.GoodsId = h_categories.GenGoodsId()
	goods2.OwnerId = "0"
	goods2.CategoryId = 0
	goods2.Avatar = "https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=3406060575,3655381718&fm=11&gp=0.jpg"
	goods2.CurrentPrice = 0
	goods2.Like = true
	goods2.CostMoney = 1
	goods2.PayRecordsCount = 0
	goods2.IsEnd = false
	goods2.TotalPrice = 9999
	goods2.Name = "iPhone XS Max 256G"
	goods2.Period = 1

	goods3.GoodsId = h_categories.GenGoodsId()
	goods3.OwnerId = "0"
	goods3.CategoryId = 0
	goods3.Avatar = "https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=3406060575,3655381718&fm=11&gp=0.jpg"
	goods3.CurrentPrice = 0
	goods3.Like = true
	goods3.CostMoney = 1
	goods3.PayRecordsCount = 0
	goods3.IsEnd = false
	goods3.TotalPrice = 9999
	goods3.Name = "iPhone XS Max 256G"
	goods3.Period = 1

	goods4.GoodsId = h_categories.GenGoodsId()
	goods4.OwnerId = "0"
	goods4.CategoryId = 0
	goods4.Avatar = "https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=3406060575,3655381718&fm=11&gp=0.jpg"
	goods4.CurrentPrice = 0
	goods4.Like = true
	goods4.CostMoney = 1
	goods4.PayRecordsCount = 0
	goods4.IsEnd = false
	goods4.TotalPrice = 9999
	goods4.Name = "iPhone XS Max 256G"
	goods4.Period = 1

	goods5.GoodsId = h_categories.GenGoodsId()
	goods5.OwnerId = "0"
	goods5.CategoryId = 0
	goods5.Avatar = "https://ss1.bdstatic.com/70cFvXSh_Q1YnxGkpoWK1HF6hhy/it/u=3406060575,3655381718&fm=11&gp=0.jpg"
	goods5.CurrentPrice = 0
	goods5.Like = true
	goods5.CostMoney = 1
	goods5.PayRecordsCount = 0
	goods5.IsEnd = false
	goods5.TotalPrice = 9999
	goods5.Name = "iPhone XS Max 256G"
	goods5.Period = 1
	o.Insert(&goods1)
	o.Insert(&goods2)
	o.Insert(&goods3)
	o.Insert(&goods4)
	o.Insert(&goods5)
}
