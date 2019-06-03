package models
import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"../conf"
)

var db *gorm.DB

//只要引用了models包，则会先调用init方法

func init() {
	var err error
	db, err = gorm.Open("mysql", conf.Sysconfig.DBUserName+":"+conf.Sysconfig.DBPassword+"@("+conf.Sysconfig.DBIp+":"+conf.Sysconfig.DBPort+")/"+conf.Sysconfig.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//如果不存在这张表，则会根据结构体创建一张表
	if !db.HasTable(&Book{}) {  //db.Set 设置一些额外的表属性                              //db.CreateTable创建表
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Book{}).Error; err != nil {
			panic(err)
		}
	}

}
