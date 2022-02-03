package main

import (
	"github.com/yamoyamoto/mecab_go/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	log.Printf("server run on port:%v", port)
	http.HandleFunc("/morphological_analysis", handlers.MorphologicalAnalysis)
	_ = http.ListenAndServe(":"+port, nil)
}
