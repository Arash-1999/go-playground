package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
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

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

func (app *Application) getSnippet(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Has("id") {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		if err != nil {
			http.NotFound(w, r)
			app.logger.Error("Not Found", "path", r.URL.Path, "snippet_id", id)
			return
		}
		app.logger.Info("get snippet", "id", id)
		fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	} else {
		stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > now() ORDER BY id DESC LIMIT 10`

		rows, err := app.db.Query(context.Background(), stmt)

		if err != nil {
			app.logger.Error("Query Failed", "error", err.Error())
		}

		defer rows.Close()

		snippets := []*Snippet{}

		for rows.Next() {
			s := &Snippet{}

			err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

			if err != nil {
				app.logger.Error("snippets", "error", err.Error())
			}

			snippets = append(snippets, s)
		}

		if err = rows.Err(); err != nil {
			app.logger.Error("snippets", "error", err.Error())
		}

		for _, snippet := range snippets {
			fmt.Fprintf(w, "%+v\n", snippet)
		}
	}
}
