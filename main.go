package main

import (
	"blog/Log"
	"log"
	"net/http"
)

func main() {
	Log.InitLog()

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
