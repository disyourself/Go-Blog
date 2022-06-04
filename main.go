package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "Godot Go Blog"
	indexData.Desc = "hello world"
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {

	var indexData IndexData
	indexData.Title = "Godot Go Blog"
	indexData.Desc = "hello world"
	t := template.New("index.html")

	//1.拿到当前路径
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexData)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
