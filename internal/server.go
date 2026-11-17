package server

import (
	"html/template"
	"net/http"
	"io/fs"
)

type App struct {
	Mux *http.ServeMux
	Tmpl *template.Template
	Sttc fs.FS
}

func NewApp(tmpl *template.Template, sttc fs.FS) *App {
	app := &App {
		Mux: http.NewServeMux(),
		Tmpl: tmpl,
		Sttc: sttc,
	}

	app.setupRoutes()
	return app
}

func (app *App) setupRoutes() {
	app.Mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(app.Sttc))))
	app.Mux.HandleFunc("GET /{$}", app.handleHome)
}
