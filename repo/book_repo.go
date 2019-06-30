package repo

import (
	"../datasource"
	"../models"
	"github.com/spf13/cast"
)
type BookRepository interface {
	List(m map[string]interface{})(total int,books []models.Book)
	Save(book models.Book)(err error)
	Get(id uint)(book models.Book,err error)
	Del(book models.Book)(err error)
}
func NewBookRepository() BookRepository{
	return &bookRepository{}
}
var db = datasource.GetDB()

type bookRepository struct {}

func (n bookRepository) List(m map[string]interface{})(total int,books []models.Book){
	db.Table("book").Count(&total)
	err := db.Limit(cast.ToInt(m["size"])).Offset((cast.ToInt(m["page"])-1)*cast.ToInt(m["size"])).Find(&books).Error
	if err!=nil {
		panic("select Error")
	}
	return
}
func (n bookRepository) Save(book models.Book)(err error){
	if book.ID != 0{
		err := db.Save(&book).Error
		return err
	}else {
		err := db.Create(&book).Error
		return err
	}
	return nil
}
func (n bookRepository) Get(id uint)(book models.Book,err error){
	err = db.First(&book,id).Error
	return
}
func (n bookRepository) Del(book models.Book)(err error){
	err = db.Unscoped().Delete(&book).Error
	return
}

