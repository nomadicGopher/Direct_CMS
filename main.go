package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
)

func main() {
	flag.Parse()

	http.HandleFunc("/", serveUserResources)

	fmt.Println("Serving user defined files...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server: ", err)
	}

	fmt.Println("User defined files served.")

	select {}
}

func serveUserResources(w http.ResponseWriter, r *http.Request) {
	filePatterns := []string{"logo.*", "sitemap.xml", "custom.css", "custom.js", "favicon.*"}

	for _, filePattern := range filePatterns {
		file, err := filepath.Glob(filepath.Join("userContent", filePattern))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error finding file pattern %s: %v", filePattern, err), http.StatusInternalServerError)
			continue
		}

		fmt.Println("Serving: 'userContent/" + file[0] + "'")

		http.ServeFile(w, r, "userContent/"+file[0])
	}
}
