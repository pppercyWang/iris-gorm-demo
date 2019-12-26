package middleware

import (
	"../models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris"
	"github.com/spf13/cast"
	"log"
	"time"
)

const JwtKey = "percy"

func GetJWT() *jwtmiddleware.Middleware {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(JwtKey), nil
		},
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		ErrorHandler: func(ctx iris.Context, s string) {

			fmt.Println("错误:", s)
			result := models.Result{Code: -1, Msg: "认证失败，请重新登录认证"}
			i, err := ctx.JSON(result)
			if err != nil {
				log.Println(i, err)
			}
		},
	})
	return jwtHandler
}

//生成token
func GenerateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nick_name": user.Username,                                            //用户信息
		"session":   user.Session,                                             //session
		"id":        user.ID,                                                  //用户信息
		"iss":       "Iris",                                                   //签发者
		"iat":       time.Now().Unix(),                                        //签发时间
		"jti":       "9527",                                                   //jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
		"exp":       time.Now().Add(10 * time.Hour * time.Duration(1)).Unix(), //过期时间
	})
	tokenString, _ := token.SignedString([]byte(JwtKey))
	fmt.Println("签发时间：", time.Now().Unix())
	fmt.Println("到期时间：", time.Now().Add(10 * time.Hour * time.Duration(1)).Unix())
	return tokenString
}

func ParseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}

func GetToken(ctx iris.Context) string {
	token := ctx.GetHeader("Authorization")
	if token != "" && len(token) > 7 {
		token = token[7:]
	}
	return token
}
func GetUserID(token string) int {
	var userId = 0
	if token != "" && token != "undefined" && len(token) > 7 {
		v, _ := ParseToken(token, JwtKey)
		if v != "" {
			userId = cast.ToInt(v.(jwt.MapClaims)["id"])
		}
	}
	return userId
}
