package general

import (
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

	w.Write([]byte("fucking home page."))
}
