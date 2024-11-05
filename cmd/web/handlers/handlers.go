package handlers

import (
	"net/http"

	"lets-go-book-2022/cmd/web/base"
	"lets-go-book-2022/cmd/web/base/middleware"
	"lets-go-book-2022/cmd/web/handlers/general"
	"lets-go-book-2022/cmd/web/handlers/snippet"
)

func Routes(app *base.Application) http.Handler {
	mux := http.NewServeMux()

	// TODO: change scope structure to pass only needed dependencies
	// snippetScope := &snippet.Snippet{
	// 	Logger: app.Logger,
	//  Db: app.Db,
	// }

	// routes
	generalScope := &general.General{
		Env: app,
	}
	mux.HandleFunc("/", generalScope.HomeHandler)
	snippetScope := &snippet.Snippet{
		Env: app,
	}
	mux.HandleFunc("/snippet", snippetScope.GetSnippet)
	mux.HandleFunc("/snippet/create", snippetScope.PostSnippet)

	// middlewares
	muMux := &MwMux{mux: mux}

	muMux.Use(middleware.SecureHeaders)

	logMiddleware := &middleware.LogMiddleware{Logger: app.Logger}
	muMux.Use(logMiddleware.LogRequest)

	recoverPanicMiddleware := &middleware.RecoverPanicMiddleware{App: app}
	muMux.Use(recoverPanicMiddleware.Recover)

	return muMux.mux
}

type MwMux struct {
	mux http.Handler
}
type Middleware func(next http.Handler) http.Handler

func (m *MwMux) Use(middleware Middleware) {
	m.mux = middleware(m.mux)
}
