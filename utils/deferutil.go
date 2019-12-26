package utils

import (
	"github.com/jinzhu/gorm"
)

func Defer(tx *gorm.DB, code *int) {
	if *code == 0{
		//提交事务
		tx.Commit()
	} else {
		//回滚
		tx.Rollback()
	}
}