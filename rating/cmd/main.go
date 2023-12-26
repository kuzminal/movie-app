package main

import (
	"log"
	"net/http"

	"github.com/kuzminal/movie-app/rating/internal/controller/rating"
	httphandler "github.com/kuzminal/movie-app/rating/internal/handler/http"
	"github.com/kuzminal/movie-app/rating/internal/repository/memory"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
