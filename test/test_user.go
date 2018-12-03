package main

import (
	`gin_project/core/h_categories`
	`gin_project/models`
	`github.com/astaxie/beego/orm`
)

func main() {
	var (
		user models.User
		ub   models.UserBalance
	)
	user.Uid = "8888888"
	user.Nick = "admin"
	user.UserDescription = "我等，生而不自由"
	user.Phone = "15913538383"
	user.Password = h_categories.GenPassword("huihui123")
	user.Avatar = "http://img4.duitang.com/uploads/item/201411/09/20141109142633_ncKBY.thumb.700_0.jpeg"
	user.Gender = 0
	user.AccountType = 0

	ub.Balance = 99999
	ub.Total = 99999
	ub.UserId = user.Uid
	orm.NewOrm().Insert(&user)
	orm.NewOrm().Insert(&ub)
}
