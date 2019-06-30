/*
@Time : 2019/5/14 11:26 
@Author : Lukebryan
@File : hmacsha256util.go
@Software: GoLand
*/
package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

func HmacSha256Encode(strings string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(strings))
	sha := hex.EncodeToString(h.Sum(nil))

	return base64.StdEncoding.EncodeToString([]byte(sha))
}
