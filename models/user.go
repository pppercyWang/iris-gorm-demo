package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Password     string
	Name         string                                         //姓名
	Email        string                         //邮箱
	Mobile       string  //手机
	QQ           string
	Gender       int                         //0男 1女
	Age          int  //年龄
	Remark       string                      //备注
	Token      string `gorm:"-"`
	Session      string `gorm:"-"`
}
