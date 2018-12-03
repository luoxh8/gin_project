package middlewares

import (
	"errors"
	`gin_project/core`
	`github.com/gin-gonic/gin`
	`net/http`
	"time"
	"github.com/dgrijalva/jwt-go"
)

/**
    JWTAuth 中间件，检查token
*/
var JWTMiddleware = core.Handler{
	"JWTAuth": func(context *gin.Context) {
		var (
			token  string
			claims *CustomClaims
			err    error
		)
		if token = context.Request.Header.Get("token"); token == "" {
			context.JSON(http.StatusBadRequest, core.ErrUserNoToken)
			context.Abort()
			return
		}
		// parseToken 解析token包含的信息
		if claims, err = NewJWT().ParseToken(token); err != nil {
			if err == TokenExpired {
				context.JSON(http.StatusBadRequest, core.ErrUserTokenExpired)
				context.Abort()
				return
			} else if err == TokenNotValidYet {
				context.JSON(http.StatusBadRequest, core.ErrUserTokenNotValidYet)
				context.Abort()
				return
			} else if err == TokenMalformed {
				context.JSON(http.StatusBadRequest, core.ErrUserTokenMalformed)
				context.Abort()
				return
			} else if err == TokenInvalid {
				context.JSON(http.StatusBadRequest, core.ErrUserTokenInvalid)
				context.Abort()
				return
			}
			context.JSON(http.StatusBadRequest, core.ErrBaseUnknown)
			context.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		context.Set("claims", claims)
	},
}

/**
    JWT 签名结构
*/
type JWT struct {
	SigningKey []byte
}

/**
    一些常量
*/
var (
	TokenExpired     = errors.New("token已经过期")
	TokenNotValidYet = errors.New("token没有激活")
	TokenMalformed   = errors.New("这不是一个token")
	TokenInvalid     = errors.New("token非法")
	SignKey          = core.SecretKey
)

/**
    载荷，可以加一些自己需要的信息
*/
type CustomClaims struct {
	Uid   string `json:"uid"`
	Nick  string `json:"name"`
	Phone string `json:"phone"`
	jwt.StandardClaims
}

/**
	新建一个jwt实例
*/
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

/**
    获取signKey
*/
func GetSignKey() string {
	return SignKey
}

/**
	设置SignKey
*/
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

/**
    CreateToken 生成一个token
*/
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

/**
    解析Token
*/
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	var (
		token *jwt.Token
		err   error
	)
	if token, err = jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	}); err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

/**
    更新token
*/
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
