package main

import (
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	fmt.Println("Start")

	http.HandleFunc("/", serveUserResources)

	if errServer := http.ListenAndServe(":8080", nil); errServer != nil {
		fmt.Println("Failed to start server: ", errServer)
	}

	select {}
}

func serveUserResources(w http.ResponseWriter, r *http.Request) {
	files := []string{"logo.*", "sitemap.xml", "custom.css", "custom.js", "favicon.*"}

	for _, pattern := range files {
		found, fileErr := filepath.Glob(filepath.Join("userContent", pattern))
		if fileErr != nil {
			http.Error(w, fmt.Sprintf("Error finding file pattern %s: %v", pattern, fileErr), http.StatusInternalServerError)
			continue
		}

		http.ServeFile(w, r, found[0])
	}
}
