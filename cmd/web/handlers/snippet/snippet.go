package snippet

import (
	"fmt"
	"lets-go-book-2022/cmd/web/base"
	"net/http"
	"strconv"
)

type Snippet struct {
	Env *base.Application
}

func (scope *Snippet) PostSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST") // it should be before WriteHeader or Write
		// w.WriteHeader(405)              // only one time per response
		// w.Write([]byte("Method Not Allowed"))
		scope.Env.Logger.Error("Method Not Allowed", "route", r.URL.Path)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	scope.Env.Logger.Info("Snippet Created", "route", r.URL.Path)
	w.Write([]byte("something created"))
}

func (scope *Snippet) GetSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		http.NotFound(w, r)
		scope.Env.Logger.Error("Not Found", "route", r.URL.Path, "id", id)
		return
	}

	scope.Env.Logger.Info("Snippet Founded", "route", r.URL.Path, "id", id)
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
