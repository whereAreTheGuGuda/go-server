package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// HmacUser 签名需要传递的参数(根据自己定义)
type HmacUser struct {
	Id    int    `json:"id"`
	Phone string `json:"phone"`
}

type MyClaims struct {
	UserId int    `json:"user_id"`
	Phone  string `json:"phone"`
	jwt.StandardClaims
}

// LoginStruct 登录的参数
type LoginStruct struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// 证书签名密钥
var jwtKey = []byte("abc")

// GenerateToken 定义生成token的方法
func GenerateToken(u HmacUser) (string, error) {
	// 定义过期时间,7天后过期
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &MyClaims{
		UserId: u.Id,
		Phone:  u.Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),     // 发布时间
			Subject:   "token",               // 主题
			Issuer:    "go-server",           // 发布者
		},
	}
	// 注意单词别写错了
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 定义解析token的方法
func ParseToken(tokenString string) (*jwt.Token, *MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
