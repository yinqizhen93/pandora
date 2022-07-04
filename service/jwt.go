package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"pandora/service/db"
	"time"
)

// Claims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type Claims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Minute * 60

const RefreshTokenExpireDuration = time.Hour * 24 * 7

var Secret = []byte("SecretKeyNotUsedInProduction")

// CreateToken 生成JWT
func CreateToken(userId int) (string, error) {
	// 创建一个我们自己的声明
	c := Claims{
		userId, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "pandora",                                  // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Secret)
}

type RefreshToken struct {
	IssuedAt    int64
	ExpiresAt   int64
	AutoRefresh bool
}

func SaveRefreshToken(ctx context.Context, id int, rt string) {
	db.Client.User.UpdateOneID(id).SetRefreshToken(rt).SaveX(ctx)
}

func CreateRefreshToken() string {
	rt := RefreshToken{
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(),
	}
	byteRt, err := json.Marshal(rt)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(byteRt)
}

//var TokenExpiredErr = errors.New("refresh token has expired")

func RefreshTokenExpired(token string) (bool, error) {
	var rt RefreshToken
	byteToken, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return false, errors.Wrap(err, "base64 DecodeString失败")
	}
	err = json.Unmarshal(byteToken, &rt)
	if err != nil {
		return false, errors.Wrap(err, "json Unmarshal失败")
	}
	if time.Unix(rt.ExpiresAt, 0).Before(time.Now()) {
		return true, nil
	}
	return false, nil
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return token, nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return token, claims, nil
	}
	return token, nil, errors.New("invalid token")
}
