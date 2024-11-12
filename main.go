package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	err error
	logFile *os.File
)

const logFileName = "appSession.log"

func main() {
	if logFile, err = os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666); err != nil {
		log.SetOutput(os.Stdout)
		log.Println("Failed to open log file: ", err)
	} else {
		log.SetOutput(logFile)
	}
	defer logFile.Close()
	log.Println(`----------------------------------------------------------------------------------------------------\n
		App session started.`)

	if err = http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Failed to start server: ", err)
	}

	http.HandleFunc("/", serveUserResources)

	select {} // Keep session open to serve functions over WASM.
}

// ---------- ACTIONS ----------
func serveUserResources(w http.ResponseWriter, r *http.Request) {
	basePath := "/user/"
	files := []string{"logo.*", "sitemap.xml", "custom.css", "custom.js", "favicon.*"}

	for _, pattern := range files {
		matches, err := filepath.Glob(filepath.Join(basePath, pattern))
		if err != nil {
			log.Printf("Error checking existence of file pattern %s: %v", pattern, err)
			continue
		}

		if len(matches) > 0 {
			http.ServeFile(w, r, matches[0])
			return
		}
	}

	log.Println("None of the custom files exist.")
	http.NotFound(w, r)
}

// ---------- UTILITIES ----------
