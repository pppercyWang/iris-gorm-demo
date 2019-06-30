package controllers

import (
	"../models"
	"../service"
	"github.com/kataras/iris"
	"github.com/spf13/cast"
)

type BookController struct {
	Ctx     iris.Context
	Service service.BookService
}
func NewBookController() *BookController {
	return &BookController{Service:service.NewBookService()}
}
func (g *BookController) PostList()(result models.Result)  {
	r := g.Ctx.Request()
	m:=make(map[string]interface{})
	m["page"] = r.PostFormValue("page")
	m["size"] = r.PostFormValue("size")
	return g.Service.List(m)
}
func (g *BookController) PostSave()(result models.Result)  {
	r := g.Ctx.Request()
	book := models.Book{}
	book.ID = cast.ToUint(r.PostFormValue("id"))
	book.Name = r.PostFormValue("name")
	book.Count =  r.PostFormValue("count")
	book.Author = r.PostFormValue("author")
	book.Type = r.PostFormValue("type")
	return g.Service.Save(book)
}
func (g *BookController) PostGet()(result models.Result)  {
	r := g.Ctx.Request()
	id := cast.ToUint(r.PostFormValue("id"))
	return g.Service.Get(id)
}
func (g *BookController) PostDel()(result models.Result)  {
	r := g.Ctx.Request()
	id := cast.ToUint(r.PostFormValue("id"))
	book:=models.Book{}
	book.ID = id
	return g.Service.Del(book)
}