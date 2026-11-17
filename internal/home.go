package server

import (
	"net/http"
)

func (app *App) handleHome(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Cherry Note",
		"PageCss": "index.css",
	}
	err := app.Tmpl.ExecuteTemplate(w, "base_app", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
}

