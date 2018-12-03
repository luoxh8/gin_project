package main

import (
	`gin_project/models`
	`github.com/astaxie/beego/orm`
)

func main() {
	o := orm.NewOrm()
	var (
		SmallIcon1 models.SmallIcon
		SmallIcon2 models.SmallIcon
		SmallIcon3 models.SmallIcon
		SmallIcon4 models.SmallIcon
		SmallIcon5 models.SmallIcon
	)

	/**
		banner
	 */

	SmallIcon1.IosActivity = "ios"
	SmallIcon1.Activity = "ios"
	SmallIcon1.Activity = "ios"
	SmallIcon1.Index = 0
	SmallIcon1.Url = "www.baidu.com"
	SmallIcon1.Title = " 类型0"
	SmallIcon1.Pic = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1540030414825&di=d30234318348dd3e613bd6c913b6dec4&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F01851855f282cf6ac7251df8d15ea0.png%401280w_1l_2o_100sh.png"
	SmallIcon1.Params = "ios"

	SmallIcon2.IosActivity = "ios"
	SmallIcon2.Activity = "ios"
	SmallIcon2.Activity = "ios"
	SmallIcon2.Index = 1
	SmallIcon2.Url = "www.baidu.com"
	SmallIcon2.Title = " 类型1"
	SmallIcon2.Pic = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1540030414825&di=d30234318348dd3e613bd6c913b6dec4&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F01851855f282cf6ac7251df8d15ea0.png%401280w_1l_2o_100sh.png"
	SmallIcon2.Params = "ios"

	SmallIcon3.IosActivity = "ios"
	SmallIcon3.Activity = "ios"
	SmallIcon3.Activity = "ios"
	SmallIcon3.Index = 2
	SmallIcon3.Url = "www.baidu.com"
	SmallIcon3.Title = " 类型2"
	SmallIcon3.Pic = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1540030414825&di=d30234318348dd3e613bd6c913b6dec4&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F01851855f282cf6ac7251df8d15ea0.png%401280w_1l_2o_100sh.png"
	SmallIcon3.Params = "ios"

	SmallIcon4.IosActivity = "ios"
	SmallIcon4.Activity = "ios"
	SmallIcon4.Activity = "ios"
	SmallIcon4.Index = 3
	SmallIcon4.Url = "www.baidu.com"
	SmallIcon4.Title = " 类型3"
	SmallIcon4.Pic = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1540030414825&di=d30234318348dd3e613bd6c913b6dec4&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F01851855f282cf6ac7251df8d15ea0.png%401280w_1l_2o_100sh.png"
	SmallIcon4.Params = "ios"

	SmallIcon5.IosActivity = "ios"
	SmallIcon5.Activity = "ios"
	SmallIcon5.Activity = "ios"
	SmallIcon5.Index = 4
	SmallIcon5.Url = "www.baidu.com"
	SmallIcon5.Title = " 类型4"
	SmallIcon5.Pic = "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1540030414825&di=d30234318348dd3e613bd6c913b6dec4&imgtype=0&src=http%3A%2F%2Fimg.zcool.cn%2Fcommunity%2F01851855f282cf6ac7251df8d15ea0.png%401280w_1l_2o_100sh.png"
	SmallIcon5.Params = "ios"

	o.Insert(&SmallIcon1)
	o.Insert(&SmallIcon2)
	o.Insert(&SmallIcon3)
	o.Insert(&SmallIcon4)
	o.Insert(&SmallIcon5)
}
