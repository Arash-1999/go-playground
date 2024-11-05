package base

import (
	"net/http"
)

// TODO: refactor global error handling
func (app *Application) ServerError(w http.ResponseWriter, err error) {
	app.Logger.Error("Internal Server Error", "error", err.Error())

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
