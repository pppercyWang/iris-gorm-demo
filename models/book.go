package models

import _"fmt"

type Book struct {
	ID        int    `gorm:"primary_key"`
	Name        string `gorm:"type:varchar(20);not null;"`
	Count        string    `gorm:"type:varchar(10);not null;"`
	Author        string    `gorm:"type:varchar(20);not null;"`
	Type     string `gorm:"type:varchar(20);not null;"`
}

func NewBook() (*Book) {
	a:= new(Book)
	return a
}

func (b *Book) Create(a *Book) error {
	return  db.Create(a).Error;
}

func (b *Book) Save(a *Book) error {
	return  	db.Save(a).Error
}

func (b *Book) Del(a *Book) (error) {
	return db.Delete(a).Error
}

func (b *Book) Get(id int) (book Book,err error) {
	err = db.First(&book, id).Error
	return
}

func (b *Book) GetList(page int,size int) (books []Book,count int,err error) {
	//如果return后面没有指定返回值，就用赋给“返回值变量”的值
	err =  db.Limit(size).Offset((page - 1) * size).Find(&books).Error
	db.Table("books").Count(&count)
	return
}