package main

import (
	"Go-Blog/common"
	"Go-Blog/router"
	"log"
	"net/http"
)

func init() {
	//预加载模板
	common.LoadTemplate()
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	//路由
	router.Router()
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
