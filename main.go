package main

import (
	"Go-Blog/config"
	"Go-Blog/models"
	"log"
	"net/http"
	"text/template"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {

	var indexData IndexData
	indexData.Title = "Godot Go Blog"
	indexData.Desc = "hello world"
	t := template.New("index.html")

	//1. 拿到当前路径
	path := config.Cfg.System.CurrentDir

	//访问博客首页模板，因为多个模板嵌套，需要对所有模板解析
	home := path + "template/home.html"
	header := path + "template/layout/header.html"
	footer := path + "template/layout/footer.html"
	personal := path + "template/layout/personal.html"
	post := path + "template/layout/post-list.html"
	pagination := path + "template/layout/pagination.html"
	t, _ = t.ParseFiles(path+"/template/index.html", home, header, footer, personal, post, pagination)

	//页面上涉及到的所有的数据，必须有定义
	Viewer := config.Viewer{
		Title:       "Godot Blog",
		Descripiton: "博客描述",
	}
	var ht = &models.HomeResponse{
		Viewer:   config.Viewer{},
		Category: []models.Category{},
		Posts:    []models.PostMore{},
		Total:    0,
		Page:     0,
		Pages:    []int{},
		PageEnd:  false,
	}

	t.Execute(w, indexData)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
