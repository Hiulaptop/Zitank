package route

import (
	"Zitank/controller"
	"Zitank/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserResetPassword struct {
	UserID      int    `json:"userid"`
	NewPassword string `json:"password"`
}

func userRouter(rs *models.AppResource) http.Handler {
	r := chi.NewRouter()
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User"))
	})
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(rs.TokenAuth))
		r.Use(jwtauth.Authenticator(rs.TokenAuth))
		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})
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
		_, tokenString, err := rs.TokenAuth.Encode(map[string]interface{}{"user_id": user.ID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(tokenString))
	})

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		var userinfo models.Users
		err := json.NewDecoder(r.Body).Decode(&userinfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userController := controller.UserController{}
		err = userController.CreateUser(rs.Store, userinfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, tokenString, err := rs.TokenAuth.Encode(map[string]interface{}{"user_id": userinfo.ID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(tokenString))
	})

	r.Post("/forgot-password", func(w http.ResponseWriter, r *http.Request) {

	})

	r.Post("/reset-password", func(w http.ResponseWriter, r *http.Request) {
		var usernpw UserResetPassword
		err := json.NewDecoder(r.Body).Decode(&usernpw)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userController := controller.UserController{}
		err = userController.ResetPassword(rs.Store, usernpw.NewPassword, usernpw.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, tokenString, err := rs.TokenAuth.Encode(map[string]interface{}{"user_id": usernpw.UserID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(tokenString))
	})
	return r
}
