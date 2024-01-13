package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type PageData struct {
	Technique string
}

func get_technique() string {
	// read source file
	content, err := os.ReadFile("techniques.txt")
	check(err)
	techniqueSlice := strings.Split(string(content), "\n")
	randomIndex := rand.Intn(len(techniqueSlice) - 1)
	slog.Debug("Index:", "randomIndex", randomIndex)
	pick := techniqueSlice[randomIndex]
	return pick
}

// check panics if an error is detected
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	tmpl := template.Must(template.ParseFiles("layout.html"))

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		technique := get_technique()
		data := PageData{
			Technique: technique,
		}
		tmpl.Execute(w, data)
		slog.Info("New response:", "technique", technique)

	}

	healthHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "OK")
		slog.Debug("Health check called")

	}

	http.HandleFunc("/random", helloHandler)
	http.HandleFunc("/health", healthHandler)
	slog.Info("Listing for requests at http://localhost:8000/random")
	http.ListenAndServe(":8000", nil)

}
