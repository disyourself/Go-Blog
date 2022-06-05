package router

import (
	"Go-Blog/api"
	"Go-Blog/views"
	"net/http"
)

func Router() {

	//1. 返回页面 views
	http.HandleFunc("/", views.HTML.Index)

	//2.数据(json)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)

	//3.静态资源
	http.Handle("/resource/", http.StripPrefix("/resouce/", http.FileServer(http.Dir("public/resouce"))))

}
