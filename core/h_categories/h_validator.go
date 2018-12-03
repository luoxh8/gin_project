package h_categories

import (
	`github.com/astaxie/beego/orm`
	`regexp`
)

var LiangHao = []string{
	"^\\d*(999)\\d*$",
	"^\\d*(888)\\d*$",
	"^\\d*(777)\\d*$",
	"^\\d*(666)\\d*$",
	"^\\d*(555)\\d*$",
	"^\\d*(444)\\d*$",
	"^\\d*(333)\\d*$",
	"^\\d*(222)\\d*$",
	"^\\d*(111)\\d*$",
	"^\\d*(000)\\d*$",
}

/**
	验证验证码
 */
func IsCode(code string) bool {
	return len(code) == 5
}

/**
	验证密码
 */
func IsPassword(password string) bool {
	regexp.MustCompile("^([A-Z]|[a-z]|[0-9]|[!@#$%^&*()_+-={}:;]){6,30}$").MatchString(password)
	return true
}

/**
	验证手机号
 */
func IsPhone(mobileNum string) bool {
	return regexp.MustCompile("^1[34578]\\d{9}$").MatchString(mobileNum)
}

/**
	判断是否是Uid
 */
func IsUid(uid string) bool {
	bytes := []byte(uid)
	if bytes[0] == '0' {
		return false
	}
	if orm.NewOrm().QueryTable("user").Filter("uid", uid).Exist() {
		return false
	}
	return true
}

func IsGoodsId(goodsId string) bool {
	if orm.NewOrm().QueryTable("goods").Filter("goods_id", goodsId).Exist() {
		return false
	}
	return true
}

/**
    是否靓号
*/
func IsGoodUid(uid string) bool {
	for _, value := range LiangHao {
		if regexp.MustCompile(value).MatchString(uid) {
			return true
		}
	}
	return false
}
