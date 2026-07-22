package main

import (
	"chitchater/handler"
	"context"
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed routes/*
var file embed.FS

func main() {
	routes, err := fs.Sub(file, "routes")
	if err != nil {
		log.Fatal(err)
	}

	setupAPI(http.FS(routes))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupAPI(routes http.FileSystem) {

	ctx := context.Background()
	serve := handler.NewChatServer(ctx)
	http.Handle("/", http.FileServer(routes))
	http.HandleFunc("/ws", serve.ServeConnection)
	http.HandleFunc("/login", serve.LoginHandler)

}
