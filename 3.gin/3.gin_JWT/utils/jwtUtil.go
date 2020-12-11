package utils

//
/*
JWT全称JSON Web Token是一种跨域认证解决方案，属于一个开放的标准，
规定了一种Token实现方式，目前多用于前后端分离项目和OAuth2.0业务场景下。
是基于Token的轻量级认证模式，服务端认证通过后，会生成一个JSON对象，
经过签名后得到一个Token（令牌）再发回给用户，
用户后续请求只需要带上这个Token，服务端解密之后就能获取该用户的相关信息了。
*/

// 我们在这里直接使用jwt-go这个库来实现我们生成JWT和解析JWT的功能。
import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

/*
我们需要定制自己的需求来决定JWT中保存哪些数据，
比如我们规定在JWT中要存储username信息，那么我们就定义一个MyClaims结构体如下：
*/

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 然后我们定义JWT的过期时间，为2小时
const TokenExpireDuration = time.Hour * 2

// 接下来还需要定义 Secret:
var MySecret = []byte("圆圆你好吗")

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username, //自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "my-project",                               //签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的 secret 签名并获得完成的编码后的字符串token
	return token.SignedString(MySecret)
}

// 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalied token")
}

