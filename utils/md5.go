/*
@Time : 2019/5/8 16:20 
@Author : Lukebryan
@File : md5.go
@Software: GoLand
*/
package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func GetMD5String(strings string) string {

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(strings))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
func Md5ByByte(bytes []byte) string {

	md5Ctx := md5.New()
	md5Ctx.Write(bytes)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMD5String(base64.URLEncoding.EncodeToString(b))
}
