package route

import (
	"net/http"

	"github.com/go-chi/jwtauth/v5"
)

func (BH BaseHandler) AdminAuthenticate(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check role
		_, claims, _ := jwtauth.FromContext(r.Context())
		userID := claims["user_id"].(int)
		role := BH.userRepositor.RoleCheck(userID)
		if role != "admin" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}
