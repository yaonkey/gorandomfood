package main

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// Структура, определяюая еду
type Food struct {
	Name        string `json:"name" sql:"name"`
	Description string `json:"description" sql:"description"`
	Image       string `json:"image" sql:"image_src"`
}

// Создание новой еды
func NewFood(name, description, image string) *Food {
	return &Food{
		Name:        name,
		Description: description,
		Image:       image,
	}
}

// Сохранение новой еды в бд
func (f *Food) Save() {
	db, err := sql.Open("sqlite3", "./foods.db")
	defer db.Close()
	if err != nil {
		log.Fatalf("Error with save: %v\n", err)
	}
}

// Преобразование патча изображения
func GetImagePath(image string) string {
	if strings.Contains(image, ".jpg") {
		image = "static/" + image
	} else {
		image = "static/" + image + ".jpg"
	}
	if _, err := os.Stat(image); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("Error with getting image path: %v\n", image)
	}
	return image
}

// Получение рандомной еды
func GetRandomFood() Food {
	db, err := sql.Open("sqlite3", "./foods.db")
	defer db.Close()

	var name string
	var description string
	var image string

	err = db.QueryRow("SELECT name, image_src, description FROM Foods ORDER BY RANDOM() LIMIT 1").Scan(&name, &image, &description)
	if err != nil {
		log.Fatalf("Error with getting random food: %v\n", err)
	}

	image = GetImagePath(image)

	food := &Food{
		Name:        name,
		Description: description,
		Image:       image,
	}
	return *food
}
