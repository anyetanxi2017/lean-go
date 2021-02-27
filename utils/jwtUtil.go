package utils

import (
	"errors"
	"gin-vue-admin/model/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

type JwtUtil struct {
}

// 常量
var (
	MySecret            = []byte("这个一个特别的冬天")
	TokenExpireDuration = time.Hour * 24 * 7 //7天过期时间
	//jwt验证成功后，放到 gin上下文中
	UserKey string = "USER_KEY"
)

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
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //过期时间
			Issuer:    "my-project",                               // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedString, err := token.SignedString(MySecret)
	if err != nil {
		panic(err)
	}
	return signedString
}

// 解析JWT
func (j JwtUtil) ParseToken(tokenStr string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
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
	claims, exists := c.Get(UserKey)
	if !exists {
		panic("获取用户信息为空")
	}
	userInfo := claims.(*JwtCustomClaims)
	return userInfo
}

// JWT 中间件
func (j JwtUtil) JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			token = c.Query("token")
			if token == "" {
				response.FailWithMessage("token为空", c)
				c.Abort()
				return
			}
		}
		jwtInfo, err := JwtUtil{}.ParseToken(token)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		c.Set(UserKey, jwtInfo)
		c.Next()
	}
}
