package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"io/fs"

	"schimith/cherry-note/internal"
)

//go:embed web/template
var myTemplate embed.FS

//go:embed web/static
var myStatic embed.FS

func main() {

	tmpl, err := template.ParseFS(myTemplate, "web/template/*/*.html")
	if err != nil {
		log.Fatalf("erro em template: %v \n", err)
	}
	
	sttc, err := fs.Sub(myStatic, "web/static")
	if err != nil {
		log.Fatalf("erro em static: %v \n", err)
	}

	app := server.NewApp(tmpl, sttc)

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", app.Mux); err != nil {
		log.Fatalf("erro no server: %v \n", err)
	}
}
