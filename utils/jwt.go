package utils

import (
	"time"

	"github.com/LightBulbClub/pole-arc/config"
	"github.com/dgrijalva/jwt-go/v4"
)

// Claims 定义了 JWT 的 payload 结构
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// getJWTSecret 从配置获取 JWT 密钥
func getJWTSecret() []byte {
	return []byte(config.GetJWTSecret())
}

// GenerateJWT 为指定用户 ID 生成 JWT
func GenerateJWT(userID uint) (string, error) {
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(time.Now().Add(time.Hour * 24 * 7).Unix())), // 7 * 24 小时过期
			Issuer:    "your-auth-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(getJWTSecret())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseJWT 解析 JWT 并返回 Claims
func ParseJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return getJWTSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
