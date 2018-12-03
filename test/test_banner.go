package main

import (
	`gin_project/models`
	`github.com/astaxie/beego/orm`
)

func main() {
	o := orm.NewOrm()
	var (
		banner1 models.Banner
		banner2 models.Banner
		banner3 models.Banner
		banner4 models.Banner
		banner5 models.Banner
	)

	/**
		banner
	 */

	banner1.IosActivity = "ios"
	banner1.Activity = "ios"
	banner1.Activity = "ios"
	banner1.Index = 0
	banner1.Url = "www.baidu.com"
	banner1.Title = " IPhone XS Max手机0"
	banner1.Pic = "http://i1.umei.cc/uploads/tu/201608/12/lbw5rld1rig.jpg"
	banner1.Params = "ios"

	banner2.IosActivity = "ios"
	banner2.Activity = "ios"
	banner2.Activity = "ios"
	banner2.Index = 1
	banner2.Url = "www.baidu.com"
	banner2.Title = " IPhone XS Max手机1"
	banner2.Pic = "http://i1.umei.cc/uploads/tu/201608/12/lbw5rld1rig.jpg"
	banner2.Params = "ios"

	banner3.IosActivity = "ios"
	banner3.Activity = "ios"
	banner3.Activity = "ios"
	banner3.Index = 2
	banner3.Url = "www.baidu.com"
	banner3.Title = " IPhone XS Max手机2"
	banner3.Pic = "http://i1.umei.cc/uploads/tu/201608/12/lbw5rld1rig.jpg"
	banner3.Params = "ios"

	banner4.IosActivity = "ios"
	banner4.Activity = "ios"
	banner4.Activity = "ios"
	banner4.Index = 3
	banner4.Url = "www.baidu.com"
	banner4.Title = " IPhone XS Max手机3"
	banner4.Pic = "http://i1.umei.cc/uploads/tu/201608/12/lbw5rld1rig.jpg"
	banner4.Params = "ios"

	banner5.IosActivity = "ios"
	banner5.Activity = "ios"
	banner5.Activity = "ios"
	banner5.Index = 4
	banner5.Url = "www.baidu.com"
	banner5.Title = " IPhone XS Max手机4"
	banner5.Pic = "http://i1.umei.cc/uploads/tu/201608/12/lbw5rld1rig.jpg"
	banner5.Params = "ios"

	o.Insert(&banner1)
	o.Insert(&banner2)
	o.Insert(&banner3)
	o.Insert(&banner4)
	o.Insert(&banner5)
}
