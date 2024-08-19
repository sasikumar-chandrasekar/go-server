package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello world! You have requested: %s\n", r.URL.Path)
		data := map[string]string{
			"message": "Hello, World!",
			"path":    r.URL.Path,
		}
		jsonResponse, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	fmt.Println("Server is running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
