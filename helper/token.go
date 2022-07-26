package helper

import "github.com/golang-jwt/jwt/v4"

type UserClaims struct {
	Identity string `json:"identity"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

var key = []byte("secret")

// GenerateToken 生成token
func GenerateToken(identity, email string) (string, error) {
	claims := &UserClaims{
		Identity:         identity,
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(key)
}

// ParseToken 解析token
func ParseToken(token string) (*UserClaims, error) {
	claims := &UserClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	},
	)
	return claims, err
}
