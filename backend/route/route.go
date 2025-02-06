package route

import (
	"Zitank/models"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ApiRouter(rs *models.AppResource) http.Handler {
	r := chi.NewRouter()
	r.Mount("/user", userRouter(rs))
	r.Mount("/music", musicRouter(rs))
	r.Mount("/post", postRouter(rs))
	r.Mount("/room", roomRouter(rs))
	return r
}
