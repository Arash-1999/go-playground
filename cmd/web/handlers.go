package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *Application) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		app.logger.Error("Not Found", "path", r.URL.Path)
		return
	}

	w.Write([]byte("fucking home page."))
}

func (app *Application) postSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST") // it should be before WriteHeader or Write
		// w.WriteHeader(405)              // only one time per response
		// w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		app.logger.Error("Method Not Allowed", "method", r.Method)
		return
	}
	w.Write([]byte("something created"))
}

func (app *Application) getSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		http.NotFound(w, r)
		app.logger.Error("Not Found", "path", r.URL.Path, "snippet_id", id)
		return
	}

	app.logger.Info("get snippet", "id", id)
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
