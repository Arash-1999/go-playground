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
}

func (scope *Snippet) GetSnippets(w http.ResponseWriter, r *http.Request) {
	snippets, err := scope.Env.Db.Snippets.Latest(r.Context())

	if err != nil {
		// scope.Env.Logger.Error("Postgres Select Error", "route", r.URL.Path, "error", err)
		scope.Env.ServerError(w, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}
}

func (scope *Snippet) GetSingleSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Invalid Snippet ID.", http.StatusBadRequest)
		return
	}

	snippet, err := scope.Env.Db.Snippets.Get(r.Context(), id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
			scope.Env.Logger.Error("Not Found", "route", r.URL.Path, "id", id)
			return
		} else {
			scope.Env.ServerError(w, err)
			fmt.Fprintf(w, "Error in getting snippet with ID %d...", id)
		}
	}

	fmt.Fprintf(w, "%+v", snippet)
}
