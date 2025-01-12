package route

import (
	"Zitank/controller"
	"Zitank/models"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func userRouter(rs *models.AppResource) http.Handler {
	r := chi.NewRouter()
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User"))
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		var userrq UserRequest
		err := json.NewDecoder(r.Body).Decode(&userrq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userController := controller.UserController{}
		user, err := userController.LoginUserByUsername(rs.Store, userrq.Username, userrq.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte(user.Username.String))
	})

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Register"))
	})

	r.Post("/logout", func(w http.ResponseWriter, r *http.Request) {

	})

	r.Post("/forgot-password", func(w http.ResponseWriter, r *http.Request) {

	})

	r.Post("/reset-password", func(w http.ResponseWriter, r *http.Request) {

	})
	return r
}
