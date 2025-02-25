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
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		users, err := BH.userRepositor.GetUsers()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		res, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(res)
	})
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(BH.TokenAuth))
		r.Use(jwtauth.Authenticator(BH.TokenAuth))
		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["userid"])))
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
		w.Write([]byte(tokenString))
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
		w.Write([]byte(tokenString))
	})
	return r
}
