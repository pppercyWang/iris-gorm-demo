/*
@Time : 2019/5/14 10:40 
@Author : Lukebryan
@File : tokenutil.go
@Software: GoLand
*/
package utils
//
//import (
//	"github.com/dgrijalva/jwt-go"
//	"time"
//)
//
//func GentToken(userID string,wechat_id string) string {
//	t := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
//		"timestamp" : time.Now().Unix(),
//		"user_id" : userID,
//		"wxid" : wechat_id,
//	})
//	token,_ := t.SignedString([]byte("secret"))
//	return "BEARER " + token
//}

//func Base64Encode(maps map[string]interface{}) string  {
//
//	b,err := json.Marshal(maps)
//	if err != nil {
//		log.Println(err)
//	}
//
//	return base64.StdEncoding.EncodeToString(b)
//}