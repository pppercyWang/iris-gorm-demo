package controllers

import (
	"github.com/kataras/iris"
	"../api"
)
//接口的入口
type BookController struct {
	Ctx     iris.Context
	Book    api.Book
}
func NewBookController() *BookController {
	return &BookController{}
}
//post :api/v1/book/save
func (g *BookController) PostSave()  {
	g.Book.Save(g.Ctx.ResponseWriter(),g.Ctx.Request())
}
//post :api/v1/book/create
func (g *BookController) PostCreate()  {
	g.Book.Create(g.Ctx.ResponseWriter(),g.Ctx.Request())
}
//post :api/v1/book/del
func (g *BookController) PostDel()  {
	g.Book.Del(g.Ctx.ResponseWriter(),g.Ctx.Request())
}
//post :api/v1/book/get
func (g *BookController) PostGet()  {
	g.Book.Get(g.Ctx.ResponseWriter(),g.Ctx.Request())
}
//post :api/v1/book/list
func (g *BookController) PostList()  {
	g.Book.GetList(g.Ctx.ResponseWriter(),g.Ctx.Request())
}