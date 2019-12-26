package service

import (
	"../models"
	"../repo"
)
type BookService interface {
	GetBookList (m map[string]interface{}) (result models.Result)
	SaveBook(book models.Book) (result models.Result)
	GetBook(id uint) (result models.Result)
	DelBook(id uint) (result models.Result)
}

type bookService struct {}

func NewBookService() BookService{
	return &bookService{}
}

var bookRepo = repo.NewBookRepository()

func (u bookService)GetBookList (m map[string]interface{}) (result models.Result){
	total,books := bookRepo.GetBookList(m)
	maps := make(map[string]interface{},2)
	maps["Total"] = total
	maps["List"] = books
	result.Data = maps
	result.Code = 0
	result.Msg ="SUCCESS"
	return
}
func (n bookService) SaveBook(book models.Book)(result models.Result){
	err := bookRepo.SaveBook(book)
	if err != nil{
		result.Code = -1
		result.Msg ="保存失败"
	}else{
		result.Code = 1
		result.Msg ="保存成功"
	}
	return
}
func (n bookService) GetBook(id uint)(result models.Result){
	book,err := bookRepo.GetBook(id)
	if err!= nil{
		result.Code = -1
		result.Msg = err.Error()
	}else{
		result.Data = book
		result.Code = 0
		result.Msg ="SUCCESS"
	}
	return
}
func (n bookService) DelBook(id uint)(result models.Result){
	err := bookRepo.DelBook(id)
	if err!= nil{
		result.Code = -1
		result.Msg = err.Error()
	}else{
		result.Code = 0
		result.Msg ="SUCCESS"
	}
	return
}

