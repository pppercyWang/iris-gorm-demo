package service

import (
	"../models"
	"../repo"
)
type BookService interface {
	List (m map[string]interface{}) (result models.Result)
	Save(book models.Book) (result models.Result)
	Get(id uint) (result models.Result)
	Del(book models.Book) (result models.Result)
}

type bookService struct {}

func NewBookService() BookService{
	return &bookService{}
}

var bookRepo = repo.NewBookRepository()

func (u bookService)List (m map[string]interface{}) (result models.Result){
	total,books := bookRepo.List(m)
	maps := make(map[string]interface{},2)
	maps["Total"] = total
	maps["List"] = books
	result.Data = maps
	result.Code = 0
	result.Msg ="SUCCESS"
	return
}
func (n bookService) Save(book models.Book)(result models.Result){
	err := bookRepo.Save(book)
	if err != nil{
		result.Data = nil
		result.Code = -1
		result.Msg ="保存失败"
	}else{
		result.Data = nil
		result.Code = 1
		result.Msg ="保存成功"
	}
	return
}
func (n bookService) Get(id uint)(result models.Result){
	book,err := bookRepo.Get(id)
	if err!= nil{
		result.Data = nil
		result.Code = -1
		result.Msg = err.Error()
	}else{
		result.Data = book
		result.Code = 0
		result.Msg ="SUCCESS"
	}
	return
}
func (n bookService) Del(book models.Book)(result models.Result){
	err := bookRepo.Del(book)
	if err!= nil{
		result.Data = nil
		result.Code = -1
		result.Msg = err.Error()
	}else{
		result.Data = nil
		result.Code = 0
		result.Msg ="SUCCESS"
	}
	return
}

