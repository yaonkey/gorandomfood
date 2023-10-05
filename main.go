package main

import (
	"fmt"
	"log"
	"net/http"
)

// Хендлер главной страницы
func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	rf := GetRandomFood()
	fmt.Fprintf(w, GetFoodView(rf))
}

// Хендлер страницы добавления контента
func addNewFood(w http.ResponseWriter, r *http.Request) {}

func main() {
	r := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/", http.StripPrefix("/static", fileServer))
	r.HandleFunc("/", mainPageHandler)
	r.HandleFunc("/add/", addNewFood)
	log.Println("Starting GoRandomFoods service...")
	log.Println("Available on http://localhost:5002/")

	go log.Fatal(http.ListenAndServe(":5002", r))
}
