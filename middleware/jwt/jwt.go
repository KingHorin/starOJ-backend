package jwt

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"starOJ-backend/config"
	"strconv"
	"strings"
	"time"
)

var Jwtkey = []byte("XingTong2568")

type MyClaims struct {
	UserID int64 `json:"userID"`
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(userID int64, expire time.Duration) (string, error) {
	claims := MyClaims{
		userID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),             // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),             // 生效时间
			Issuer:    "starOJ",                                   // 颁发者签名
		},
	}
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenStruct.SignedString(Jwtkey)
}

// VerifyToken 验证token
func VerifyToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Jwtkey, nil
	})

	if err != nil {
		return nil, errors.New("解析token失败")
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token不合法")
}

// AccessToken 获取token
func AccessToken(userId int64) (string, error) {
	return GenerateToken(userId, 2*time.Hour)
}

// RefreshToken 刷新token，即重新颁发一个
func RefreshToken(userId int64) (string, error) {
	return GenerateToken(userId, 5*time.Hour)
}

// JwtMiddleware jwt中间件
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")

		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 0, "msg": "未携带token"})
			c.Abort()
			return
		}
		tokenSlice := strings.SplitN(tokenStr, " ", 2)
		if len(tokenSlice) != 2 && tokenSlice[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 0, "msg": "token格式错误"})
			c.Abort()
			return
		}
		tokenStruck, err := VerifyToken(tokenSlice[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": "0", "msg": "token验证失败"})
			c.Abort()
			return
		}
		tokenInRedis, err := config.GetRD().HGet("token", strconv.FormatInt(tokenStruck.UserID, 10)).Result()
		if err != nil || tokenInRedis != tokenSlice[1] {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 0, "msg": "token不正确"})
			c.Abort()
			return
		}

		c.Set("userID", tokenStruck.UserID)
		c.Next()
	}
}
