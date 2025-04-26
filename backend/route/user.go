package route

import (
	"Zitank/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

type UserRegister struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
}

func FromUserRegister(ur UserRegister) models.Users {
	return models.Users{
		Username:    strings.ToLower(ur.Username),
		Password:    strings.ToLower(ur.Password),
		Email:       strings.ToLower(ur.Email),
		Fullname:    strings.ToLower(ur.Fullname),
		PhoneNumber: strings.ToLower(ur.PhoneNumber),
		Gender:      "None",
		Role:        "user",
	}
}

func (BH BaseHandler) userRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(jwtauth.Verifier(BH.TokenAuth))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		type GetBody struct {
			Status string          `json:"status"`
			Users  []*models.Users `json:"users"`
		}
		_, claims, _ := jwtauth.FromContext(r.Context())
		userID, ok := claims["userid"]
		if !ok || BH.userRepositor.RoleCheck(int(userID.(float64))) != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 - Unauthorized"))
			return
		}
		users, err := BH.userRepositor.GetUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		var body GetBody
		body.Status = "success"
		body.Users = users
		res, err := json.Marshal(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(res)
	})
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Authenticator(BH.TokenAuth))
		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())

			userID, ok := claims["userid"].(int)
			role := BH.userRepositor.RoleCheck(userID)
			if !ok || role != "admin" {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("false"))
				return
			}

			w.Write([]byte(role))
		})
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		var userrq UserRequest
		err := json.NewDecoder(r.Body).Decode(&userrq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := BH.userRepositor.LoginUserByUsername(strings.ToLower(userrq.Username), strings.ToLower(userrq.Password))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, tokenString, err := BH.TokenAuth.Encode(map[string]interface{}{"userid": user.ID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"token":  tokenString,
		})
	})

	r.Post("/register", func(w http.ResponseWriter, r *http.Request) {
		var userinfo UserRegister
		err := json.NewDecoder(r.Body).Decode(&userinfo)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userObject := FromUserRegister(userinfo)

		id, err := BH.userRepositor.RegisterUser(&userObject)
		if err != nil {
			fmt.Printf("CreateUser error: %v\n", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, tokenString, err := BH.TokenAuth.Encode(map[string]interface{}{"userid": id})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"token":  tokenString,
		})
	})

	r.Post("/forgot-password", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hihi"))
	})

	r.Post("/reset-password", func(w http.ResponseWriter, r *http.Request) {
		var usernpw UserResetPassword
		err := json.NewDecoder(r.Body).Decode(&usernpw)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = BH.userRepositor.ResetPassword(usernpw.NewPassword, usernpw.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, tokenString, err := BH.TokenAuth.Encode(map[string]interface{}{"userid": usernpw.UserID})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"token":  tokenString,
		})
	})
	return r
}
