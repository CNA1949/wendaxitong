package util

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"wendaxitong/api_gin_gateway/pkg/codeMsg"
)

// UserClaims 用户信息类，作为生成token的参数
type UserClaims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//func (uc *UserClaims) Valid() error { return nil }

/* 自定义token密钥 */
var (
	AccessSecret  = []byte("accessSecret")
	RefreshSecret = []byte("refreshSecret")
)

const (
	AccessTokenKeySuffix      = "_accessToken"
	RefreshTokenKeySuffix     = "_refreshToken"
	AccessTokenKeyExpireTime  = 300  // AccessToken 过期时间5分钟
	RefreshTokenKeyExpireTime = 3600 //RefreshToken 过期时间1小时
)

// GenerateToken 生成token
func GenerateToken(userName, password string, secret []byte) (string, error) {
	claims := UserClaims{
		userName,
		password,
		jwt.StandardClaims{
			Issuer:   "WenDaXiTong",
			IssuedAt: time.Now().Unix(),
		},
	}

	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	// 加密token
	tokenSigned, err := token.SignedString(secret)
	if err != nil {
		fmt.Println("获取token失败，Secret错误,err:", err)
		return "", err
	}

	return tokenSigned, nil
}

// GetARToken 获取accessToken和refreshToken
func GetARToken(userName, password string) (string, string, error) {
	// 获取accessToken
	accessTokenSigned, err := GenerateToken(userName, password, AccessSecret)
	if err != nil {
		fmt.Println("获取accessToken失败，err:", err)
		return "", "", err
	}
	// 获取refreshToken
	refreshTokenSigned, err := GenerateToken(userName, password, RefreshSecret)
	if err != nil {
		fmt.Println("获取refreshToken失败，err:", err)
		return "", "", err
	}

	// 将token存入redis
	err = SetKeyValue(userName+AccessTokenKeySuffix, accessTokenSigned, AccessTokenKeyExpireTime)
	if err != nil {
		fmt.Println("err:", err.Error())
		return "", "", err
	}
	err = SetKeyValue(userName+RefreshTokenKeySuffix, refreshTokenSigned, RefreshTokenKeyExpireTime)
	if err != nil {
		fmt.Println("err:", err.Error())
		return "", "", err
	}

	return accessTokenSigned, refreshTokenSigned, nil
}

// ParseToken 解密token
func ParseToken(tokenString string, secret []byte) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, errors.New("token格式错误")
	}
	if claims, ok := token.Claims.(*UserClaims); ok {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// RefreshAccessToken 根据refreshToken刷新accessToken
func RefreshAccessToken(userName string) (string, uint64, error) {
	rToken, err := GetValueByKey(userName + RefreshTokenKeySuffix)
	if err != nil {
		fmt.Println("GetValueByKey() err:", err.Error())
		return "", codeMsg.ErrorInvalidToken, err
	}
	parseToken, err := ParseToken(rToken, RefreshSecret)
	if err != nil {
		fmt.Println("ParseToken() err:", err.Error())
		return "", codeMsg.Failed, err
	}

	// 刷新accessToken
	accessTokenSigned, err := GenerateToken(userName, parseToken.Password, AccessSecret)
	if err != nil {
		fmt.Println("GenerateToken()，err:", err)
		return "", codeMsg.Failed, err
	}
	// 将accessTokenSigned存入redis
	err = SetKeyValue(userName+AccessTokenKeySuffix, accessTokenSigned, 300)
	if err != nil {
		fmt.Println("SetKeyValue():刷新时accessToken存入redis失败，err:", err)
		return "", codeMsg.Failed, err
	}

	return accessTokenSigned, codeMsg.SUCCESS, nil
}

// DeleteARToken 删除accessToken 和 refreshToken
func DeleteARToken(userName string) error {
	err := DeleteKeyValue(userName + AccessTokenKeySuffix)
	if err != nil {
		return err
	}
	err = DeleteKeyValue(userName + RefreshTokenKeySuffix)
	if err != nil {
		return err
	}
	return nil
}
