package handlers

import (
	"net/http"

	"lets-go-book-2022/cmd/web/base"
	"lets-go-book-2022/cmd/web/base/middleware"
	"lets-go-book-2022/cmd/web/handlers/general"
	"lets-go-book-2022/cmd/web/handlers/snippet"
	"lets-go-book-2022/cmd/web/handlers/user"
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
	mux.HandleFunc("GET /snippet", snippetScope.GetSnippets)
	mux.HandleFunc("POST /snippet", snippetScope.PostSnippet)
	mux.HandleFunc("GET /snippet/{id}", snippetScope.GetSingleSnippet)

	userScope := &user.User{
		Env: app,
	}
	mux.HandleFunc("POST /user", userScope.PostUser)

	// middlewares
	mwMux := &MwMux{mux: mux}

	mwMux.Use(middleware.SecureHeaders)

	logMiddleware := &middleware.LogMiddleware{Logger: app.Logger}
	mwMux.Use(logMiddleware.LogRequest)

	recoverPanicMiddleware := &middleware.RecoverPanicMiddleware{App: app}
	mwMux.Use(recoverPanicMiddleware.Recover)

	return mwMux.mux
}

type MwMux struct {
	mux http.Handler
}
type Middleware func(next http.Handler) http.Handler

func (m *MwMux) Use(middleware Middleware) {
	m.mux = middleware(m.mux)
}
