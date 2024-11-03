package handlers

import (
	"net/http"

	"lets-go-book-2022/cmd/web/base"
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
	snippetScope := &snippet.Snippet{
		Env: app,
	}
	generalScope := &general.General{
		Env: app,
	}
	mux.HandleFunc("/", generalScope.HomeHandler)

	mux.HandleFunc("/snippet", snippetScope.GetSnippet)
	mux.HandleFunc("/snippet/create", snippetScope.PostSnippet)

	return mux
}
