package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 创建一个struct 用来生成token 具体的用法可以参照jwt-go
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// 创建用户签名的JWT key
// 关于jwtKey的生成可以使用自己的算法来保证唯一性，这里为了简单就直接用字符串代替了
var jwtKey = []byte("my_secret_key")
// 创建token
func CreateToken(email string, expirationTime time.Time) (tokenString string, err error) {
	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	claims.Email = email
	// 声明生成token的算法和claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 创建token字符串
	return token.SignedString(jwtKey)
}

// 校验token是否有效
func CheckToken(tknStr string) (b bool, t *jwt.Token, c *Claims) {
	tkn, err := jwt.ParseWithClaims(tknStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false, nil, nil
	}
	_claims, _ := tkn.Claims.(*Claims)
	if !tkn.Valid {
		return false, nil, _claims
	}
	return true, tkn, _claims
}
