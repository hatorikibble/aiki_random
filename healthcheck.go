package main

import (
	"net/http"
	"os"
)

func main() {
	_, err := http.Get("http://127.0.0.1:8000/health")
	if err != nil {
		os.Exit(1)
	}
}
