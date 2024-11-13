package user

import (
	"fmt"
	"lets-go-book-2022/cmd/web/base"
	"lets-go-book-2022/internal/data/models"
	"net/http"
)

type User struct {
	Env *base.Application
}

// TODO: manage json structs folder structre
type PostUserBody struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (scope *User) PostUser(w http.ResponseWriter, r *http.Request) {
	var input PostUserBody
	err := scope.Env.ReadJSON(w, r, &input)

	// TODO: add schema validation

	fmt.Println(err)
	if err != nil {
		scope.Env.ClientError(w, http.StatusBadRequest)
		return
	}

	user := &models.User{
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.Password.Set(input.Password)

	if err != nil {
		scope.Env.ServerError(w, err)
	}

	err = scope.Env.Db.User.Insert(r.Context(), user)

	if err != nil {
		scope.Env.Logger.Error("Postgres Insert Error", "route", r.URL.Path, "error", err, "data", input)
		return
	}

	scope.Env.Logger.Info("Snippet Created", "route", r.URL.Path)
}
