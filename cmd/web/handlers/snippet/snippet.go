package snippet

import (
	"encoding/json"
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

// TODO: manage json structs folder structre
type PostSnippetBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (scope *Snippet) PostSnippet(w http.ResponseWriter, r *http.Request) {
	var snippet PostSnippetBody
	json.NewDecoder(r.Body).Decode(&snippet)

	if snippet.Title == "" {
		scope.Env.ClientError(w, http.StatusBadRequest)
		return
	}
	if snippet.Content == "" {
		scope.Env.ClientError(w, http.StatusBadRequest)
		return
	}

	id, err := scope.Env.Db.Snippets.Insert(r.Context(), snippet.Title, snippet.Content)

	if err != nil {
		scope.Env.Logger.Error("Postgres Insert Error", "route", r.URL.Path, "error", err, "data", snippet)
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
