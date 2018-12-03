package models

import (
	`gin_project/core`
	`github.com/astaxie/beego/orm`
	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/mattn/go-sqlite3"
	`time`
)

type Base struct {
	Id      int       `orm:"auto" json:"-"`
	Created time.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	Updated time.Time `orm:"auto_now;type(datetime)" json:"-"`
}

func init() {
	initMysql()
}

func initMysql() {
	orm.Debug = core.Debug
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", core.Mysql)
	/**
		模型同步
	 */
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(UserBalance))
	orm.RegisterModel(new(Category))
	orm.RegisterModel(new(Banner))
	orm.RegisterModel(new(PayMan))
	orm.RegisterModel(new(SmallIcon))
	orm.RegisterModel(new(Goods))
	orm.RegisterModel(new(GoodsBannerPics))
	orm.RegisterModel(new(HomeGoods))
	orm.RegisterModel(new(GoodsPayMans))
	orm.RegisterModel(new(RealTimeAuction))
	orm.RegisterModel(new(Collection))
	orm.RegisterModel(new(GoodsRoomCount))
	if err := orm.RunSyncdb("default", core.OrmForce, core.OrmVerbose); err != nil {
		panic(err.Error())
	}
}
