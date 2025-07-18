package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// Claims 定义了 JWT 的 payload 结构
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // 从环境变量获取 JWT 密钥

// GenerateJWT 为指定用户 ID 生成 JWT
func GenerateJWT(userID uint) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(time.Now().Add(time.Hour * 24).Unix())), // 24 小时过期
			Issuer:    "your-auth-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseJWT 解析 JWT 并返回 Claims
func ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
