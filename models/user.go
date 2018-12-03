package models

import (
	`gin_project/core/h_categories`
	`github.com/astaxie/beego/orm`
)

/**
	用户表
 */
type User struct {
	Base
	Uid             string `orm:"size(7);unique;index" json:"uid"`    // 用户id
	Nick            string `orm:"size(50)" json:"nick"`               // 用户名
	UserDescription string `orm:"type(text)" json:"user_description"` // 用户描述
	Phone           string `orm:"size(20)" json:"phone"`              // 手机号
	Password        string `orm:"size(128)" json:"-"`                 // 密码
	Avatar          string `orm:"size(1024)" json:"avatar"`           // 用户头像
	Gender          int    `orm:"default(1)" json:"gender"`           // 用户性别 0.男 1.女
	AccountType     int    `orm:"default(1)" json:"account_type"`     // 用户类型 0.官方账号 1.用户 2.机器人
	Lock            bool   `orm:"default(false)" json:"lock"`         // 是否锁定
	OauthOpenId     string `orm:"size(128)" json:"oauth_open_id"`     // 第三方OpenID
	OauthAvatar     string `orm:"size(200)" json:"oauth_avatar"`      // 第三方头像
	OauthForm       string `orm:"size(20)" json:"oauth_form"`         // 微博 微信 qq
}

func CreateNormalUser(phone, password, avatar string) *User {
	user := User{}
	user.Uid = h_categories.GenUid()
	user.Nick = "手机用户" + phone
	user.Phone = phone
	user.Password = h_categories.GenPassword(password)
	user.Avatar = avatar
	user.Gender = 1
	user.AccountType = 1
	user.Lock = false
	return &user
}

func (user *User) FindByPhone(phone string) bool {
	var (
		err error
	)
	if err = orm.NewOrm().QueryTable("user").Filter("phone", phone).One(user); err != nil {
		return false
	}
	return true
}

func (user *User) IsLoginParamsPass(phone, password string) bool {
	if !h_categories.IsPassword(password) || !h_categories.IsPhone(phone) {
		return false
	}
	return true
}

func (user *User) IsRegisterParamsPass(phone, password, code string) bool {
	if !h_categories.IsPassword(password) || !h_categories.IsPhone(phone) || code != "66666" {
		return false
	}
	return true
}

func (user *User) IsPasswordCorrect(password string) bool {
	if user.Password != h_categories.Reverse(h_categories.GenMD5String(password)) {
		return false
	}
	return true
}

/**
    用户余额
*/
type UserBalance struct {
	Base
	UserId  string `orm:"size(7);unique;index" json:"user_id"` // user外键
	Balance int64  `json:"balance"`                            // 竞拍当前金币
	Total   int64  `json:"total"`                              // 累计充值
}

func (ub *UserBalance) FindByUserId(uid string) bool {
	if err := orm.NewOrm().QueryTable("user_balance").Filter("user_id", uid).One(ub); err != nil {
		return false
	}
	return true
}

func (ub *UserBalance) SaveBalance() bool {
	if _, err := orm.NewOrm().Update(ub.Balance); err != nil {
		return false
	}
	return true
}

func CreateNormalUserBalance(user *User) *UserBalance {
	balance := UserBalance{}
	balance.UserId = user.Uid
	balance.Balance = 0
	balance.Total = 0
	return &balance
}

/**
    用户收藏
*/
type Collection struct {
	Base
	UserId  string `orm:"size(7);unique;index" json:"user_id"`
	GoodsId string `orm:"size(64);index;unique" json:"goods_id"`
}

func CreateNormalCollection(user *User, goods *Goods) *Collection {
	return &Collection{
		UserId:  user.Uid,
		GoodsId: goods.GoodsId,
	}
}
