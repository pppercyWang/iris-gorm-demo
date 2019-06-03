package utils

import (
	"github.com/jinzhu/gorm"
)

func Defer(tx *gorm.DB, flag *bool) {
	if *flag{
		//提交事务
		tx.Commit()
	} else {
		//回滚
		tx.Rollback()
	}
}