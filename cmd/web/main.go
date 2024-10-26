package main

import (
	"flag"
	"fmt"
	"lets-go-book-2022/internal/config"
	"log"
	"net/http"
)

var configPath = flag.String("config", "internal/config/config.yaml", "yaml config path")

func main() {
	config.Load(*configPath)

	// TODO: config logger

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)

	mux.HandleFunc("/snippet", getSnippet)
	mux.HandleFunc("/snippet/create", postSnippet)

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", config.Configs.Server.Port),
		// TODO: pass error logger
		// ErrorLog: ,
		Handler: mux,
	}

	log.Println("Starting server on :", config.Configs.Server.Port)
	err := server.ListenAndServe()
	log.Fatal(err)
}
