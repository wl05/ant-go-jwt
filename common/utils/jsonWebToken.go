package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	email string `json:"email"`
	jwt.StandardClaims
}

// 创建用户签名的JWT key
var jwtKey = []byte("my_secret_key") // 关于jwtKey的生成可以使用自己的算法来保证唯一性，这里为了简单就直接用字符串代替了
// generate token
func CreateToken(email string, expirationTime time.Time) (tokenString string, err error) {
	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	//expirationTime := time.Now().Add(5 * time.Minute).Unix()
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		email: email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	return token.SignedString(jwtKey)
}

// 校验token是否有效
func CheckToken(tknStr string) (b bool, t *jwt.Token) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid {
		return false, nil
	}
	if err != nil {
		return false, nil
	}
	return true, tkn
}
