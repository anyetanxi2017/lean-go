package utils

import (
	"errors"
	"gin-vue-admin/constant"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type JwtUtil struct {
}
type JwtCustomClaims struct {
	ID       uint
	Username string
	jwt.StandardClaims
}

// 生成token
func (j JwtUtil) GenToken(uid uint, username string) string {
	c := JwtCustomClaims{
		ID:       uid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constant.TokenExpireDuration).Unix(), //过期时间
			Issuer:    "my-project",                                        // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedString, err := token.SignedString(constant.MySecret)
	if err != nil {
		panic(err)
	}
	return signedString
}

// 解析JWT
func (j JwtUtil) ParseToken(tokenStr string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return constant.MySecret, nil
	})
	if err != nil {
		return nil, errors.New("token已过期，请重新登录")
	}
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token已过期，请重新登录")
}

func (j JwtUtil) GetClaims(c *gin.Context) *JwtCustomClaims {
	claims, exists := c.Get(constant.UserKey)
	if !exists {
		panic("获取用户信息为空")
	}
	userInfo := claims.(*JwtCustomClaims)
	return userInfo
}
