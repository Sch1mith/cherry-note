package main

import (
	"embed"
	"fmt"
	"io/fs"
	"html/template"
	"log"
	"net/http"
	"os"
)

//go:embed all:web/template
var templateFS embed.FS

//go:embed all:web/static
var staticFS embed.FS

func amain() {
	// 1. Parse all HTML files across pages, layouts, and components recursively
	tmpl, err := template.ParseFS(templateFS, "web/template/*/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}

	// 2. Initialize the router
	mux := http.NewServeMux()

	// 3. Serve static files (CSS, JS) from the embedded static folder using io/fs
	staticSub, err := fs.Sub(staticFS, "web/static")
	if err != nil {
		log.Fatalf("Failed to create static sub-filesystem: %v", err)
	}
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticSub))))

	// 4. Define your home route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		data := map[string]interface{}{
			"Title":   "Cherry Note",
			"PageCSS": "index.css", // Links to web/static/css/index.css
		}

		err := tmpl.ExecuteTemplate(w, "base_app", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// 5. Start the local server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Server stopped: %v", err)
	}
}
