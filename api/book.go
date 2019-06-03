package api

import (
	"../models"
	"../utils"
	"net/http"
	"strconv"
)

type Book struct {
}

func (b Book) Create(w http.ResponseWriter, r *http.Request){
	name := r.PostFormValue("name")
	count := r.PostFormValue("count")
	types := r.PostFormValue("type")
	author := r.PostFormValue("author")
	m := models.NewBook()
	m.Name = name
	m.Count = count
	m.Type = types
	m.Author = author
	err := m.Create(m)
	if err != nil {
		rel, _ := utils.JsonEncode(-1, nil, err.Error())
		w.Write(rel)
		return
	} else {
		rel, _ := utils.JsonEncode(0, m.ID, "创建成功")
		w.Write(rel)
		return
	}
}
func (b Book) Save(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	name := r.PostFormValue("name")
	count := r.PostFormValue("count")
	types := r.PostFormValue("type")
	author := r.PostFormValue("author")
	m := models.NewBook()
	id2, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	m.ID = int(id2)
	m.Name = name
	m.Count = count
	m.Type = types
	m.Author = author
	err = m.Save(m)
	if err != nil {
		rel, _ := utils.JsonEncode(-1, nil, err.Error())
		w.Write(rel)
		return
	} else {
		rel, _ := utils.JsonEncode(0, m.ID, "修改成功")
		w.Write(rel)
		return
	}
}

func (b Book) Del(w http.ResponseWriter, r *http.Request) {
	id := r.PostFormValue("id")
	m := models.NewBook()
	id2, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	m.ID = int(id2)
	err = m.Del(m)
	if err != nil {
		rel, _ := utils.JsonEncode(-1, nil, err.Error())
		w.Write(rel)
		return
	} else {
		rel, _ := utils.JsonEncode(0, m.ID, "删除成功")
		w.Write(rel)
		return
	}
}

func (b Book) Get(w http.ResponseWriter, r *http.Request) {
	m := models.NewBook()
	id1 := r.PostFormValue("id")
	id, _ := strconv.ParseInt(id1, 10, 64)
	book,err := m.Get(int(id))
	if err != nil {
		rel, _ := utils.JsonEncode(-1, nil, err.Error())
		w.Write(rel)
		return
	} else {
		rel, _ := utils.JsonEncode(0, book, "查询成功")
		w.Write(rel)
		return
	}
}

func (b Book) GetList(w http.ResponseWriter, r *http.Request) {
	m := models.NewBook()
	page1 := r.PostFormValue("page")
	page, _ := strconv.ParseInt(page1, 10, 64)
	size1 := r.PostFormValue("size")
	size,_ := strconv.ParseInt(size1,10,64)
	books,count,err := m.GetList(int(page), int(size))
	dict := map[string]interface{}{"list": books, "total": count}
	if err != nil {
		rel, _ := utils.JsonEncode(-1, nil, err.Error())
		w.Write(rel)
		return
	} else {
		rel, _ := utils.JsonEncode(0, dict, "查询成功")
		w.Write(rel)
		return
	}
}