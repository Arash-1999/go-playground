package general

import (
	"fmt"
	"lets-go-book-2022/cmd/web/base"
	"net/http"
)

type General struct {
	Env *base.Application
}

func (scope *General) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		scope.Env.Logger.Error("Not Found", "route", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	snippets, err := scope.Env.Db.Snippets.Latest(r.Context())

	if err != nil {
		scope.Env.Logger.Error("Postgres Select Error", "route", r.URL.Path, "error", err)
		http.NotFound(w, r)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}
}
