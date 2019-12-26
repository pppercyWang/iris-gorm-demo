
package repo

import (
	"../datasource"
	"../models"
	"../utils"
	// "github.com/spf13/cast"
	"log"
)

type UserRepository interface {
	GetUserByUserNameAndPwd(username string, password string) (user models.User)
	GetUserByUsername(username string) (user models.User)
	Save(user models.User) (int, models.User)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct{}

//登录
func (n userRepository) GetUserByUserNameAndPwd(username string, password string) (user models.User) {
	db := datasource.GetDB()
	db.Where("username = ? and password = ?", username, password).First(&user)
	return
}
func (n userRepository) GetUserByUsername(username string) (user models.User) {
	db := datasource.GetDB()
	db.Where("username = ?", username).First(&user)
	return
}

//添加/修改
func (n userRepository) Save(user models.User) (int, models.User) {
	code := 0
	tx := datasource.GetDB().Begin()
	defer utils.Defer(tx,&code)
	if user.ID != 0 {
		var oldUser models.User
		datasource.GetDB().First(&oldUser, user.ID)
		user.CreatedAt = oldUser.CreatedAt
		user.Username = oldUser.Username
		if user.Name == "" {
			user.Name = oldUser.Name
		}
		if user.Email == "" {
			user.Email = oldUser.Email
		}
		if user.Mobile == "" {
			user.Mobile = oldUser.Mobile
		}
		if user.QQ == "" {
			user.QQ = oldUser.QQ
		}
		if user.Gender == 0 {
			user.Gender = oldUser.Gender
		}
		if user.Age == 0 {
			user.Age = oldUser.Age
		}
		if user.Remark == "" {
			user.Remark = oldUser.Remark
		}
	}
	if user.Password != "" {
		user.Password = utils.GetMD5String(user.Password)
	}
	if err := tx.Save(&user).Error; err != nil {
		log.Println(err)
		code = -1
	}
	return code, user
}
