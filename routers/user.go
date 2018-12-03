package routers

import (
	`gin_project/core`
	`gin_project/middlewares`
	`gin_project/models`
	`github.com/astaxie/beego/orm`
	"github.com/gin-gonic/gin"
	`net/http`
)

type UsersLoginForm struct {
	Phone    string `form:"phone" json:"phone" xml:"phone" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type UsersRegisterForm struct {
	Phone    string `form:"phone" json:"phone" xml:"phone" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
	Code     string `form:"code" json:"code" xml:"code" binding:"required"`
}

var UsersRouter = core.Handler{
	/**
		登录
	 */
	"login": func(context *gin.Context) {
		var (
			login UsersLoginForm
			user  models.User
			token string
			jwt   middlewares.JWT
			err   error
		)

		/**
			参数校验
		 */
		if err = context.ShouldBind(&login); err != nil {
			context.SecureJSON(http.StatusBadRequest, core.ErrBaseParams)
			return
		}
		if !user.IsLoginParamsPass(login.Phone, login.Password) {
			context.SecureJSON(http.StatusBadRequest, core.ErrBaseParams)
			return
		}
		/**
			查找用户
		 */
		if !user.FindByPhone(login.Phone) {
			context.SecureJSON(http.StatusBadRequest, core.ErrUserDoesNotExist)
			return
		}
		/**
			密码校验
		 */
		if !user.IsPasswordCorrect(login.Password) {
			context.SecureJSON(http.StatusBadRequest, core.ErrUserPassword)
			return
		}
		/**
			登录成功
		 */
		jwt = middlewares.JWT{SigningKey: []byte(core.SecretKey)}
		claims := middlewares.CustomClaims{
			Uid:            user.Uid,
			Phone:          user.Phone,
			Nick:           user.Nick,
		}
		if token, err = jwt.CreateToken(claims); err != nil {
			context.SecureJSON(http.StatusBadRequest, core.ErrUserTokenCreate)
			return
		}
		context.SecureJSON(http.StatusOK, gin.H{
			"code": core.SuccessCode,
			"msg":  "登录成功",
			"data": gin.H{
				"user":  user,
				"token": token,
			},
		})
	},

	/**
	    注册
	*/
	"register": func(context *gin.Context) {
		var (
			register UsersRegisterForm
			user     models.User
			token    string
			jwt      middlewares.JWT
			claims   middlewares.CustomClaims
			err      error
		)

		/**
			参数校验
		 */
		if err = context.ShouldBind(&register); err != nil {
			context.SecureJSON(http.StatusBadRequest, core.ErrBaseParams)
			return
		}

		if !user.IsRegisterParamsPass(register.Phone, register.Password, register.Code) {
			context.SecureJSON(http.StatusBadRequest, core.ErrBaseParams)
			return
		}

		/**
			查找用户
		 */
		if user.FindByPhone(register.Phone) {
			context.SecureJSON(http.StatusBadRequest, core.ErrUserAlreadyExists)
			return
		}

		/**
			是否违规注册
		 */

		/**
			注册成功
		 */
		tempUser := models.CreateNormalUser(register.Phone, register.Password, "")
		balance := models.CreateNormalUserBalance(tempUser)
		jwt = middlewares.JWT{SigningKey: []byte(core.SecretKey)}
		claims = middlewares.CustomClaims{
			Uid:   tempUser.Uid,
			Phone: tempUser.Phone,
			Nick:  tempUser.Nick,
		}
		if token, err = jwt.CreateToken(claims); err != nil {
			context.SecureJSON(http.StatusBadRequest, core.ErrUserTokenCreate)
			return
		}
		o := orm.NewOrm()
		o.Insert(tempUser)
		o.Insert(balance)
		context.SecureJSON(http.StatusOK, gin.H{
			"code": core.SuccessCode,
			"msg":  "注册成功",
			"data": gin.H{
				"user":         tempUser,
				"user_balance": balance,
				"token":        token,
			},
		})
	},
	/**
		第三方登录
	*/
	"thirdLogin": func(context *gin.Context) {

	},
	/**
	    用户信息
	*/
	"info": func(context *gin.Context) {
		value, _ := context.Get("claims")
		claims := value.(*middlewares.CustomClaims)
		user := models.User{Uid: claims.Uid}
		balance := models.UserBalance{UserId: claims.Uid}
		orm.NewOrm().Read(&user)
		context.SecureJSON(200, gin.H{
			"code": core.SuccessCode,
			"data": gin.H{
				"user":         user,
				"user_balance": balance,
			},
		})
	},
	/**
	    验证码
	*/
	"sendCode": func(context *gin.Context) {
		context.SecureJSON(200, core.SuccessCodeSend)
	},
	/**
	    用户充值记录
	*/
	"rechargeList": func(context *gin.Context) {

	},
}
