package snippet

import (
	"errors"
	"fmt"
	"lets-go-book-2022/cmd/web/base"
	"lets-go-book-2022/internal/data/models"
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

	// TODO: get title and content from request body
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut Slowly, slowly!\n\n- Kobayashi Issa"

	id, err := scope.Env.Db.Snippets.Insert(r.Context(), title, content)

	if err != nil {
		scope.Env.Logger.Error("Postgres Insert Error", "route", r.URL.Path, "error", err, "title", title, "content", content)
		return
	}

	scope.Env.Logger.Info("Snippet Created", "route", r.URL.Path)
	fmt.Fprintf(w, "Snippet created with id: %d", id)
	// http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
	// w.Write([]byte("something created"))
}

func (scope *Snippet) GetSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		http.NotFound(w, r)
		scope.Env.Logger.Error("Not Found", "route", r.URL.Path, "id", id)
		return
	}

	snippet, err := scope.Env.Db.Snippets.Get(r.Context(), id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
			scope.Env.Logger.Error("Not Found", "route", r.URL.Path, "id", id)
			return
		} else {
			fmt.Fprintf(w, "Error in getting snippet with ID %d...\n %s", id, err)
		}
	}

	scope.Env.Logger.Info("Snippet Founded", "route", r.URL.Path, "id", id, "data", snippet)
	// fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	fmt.Fprintf(w, "%+v", snippet)
}
