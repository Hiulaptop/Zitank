package route

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (BH BaseHandler) ApiRouter() http.Handler {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User"))
	})
	r.Mount("/user", BH.userRouter())
	r.Mount("/music", BH.musicRouter())
	r.Mount("/post", BH.postRouter())
	r.Mount("/room", BH.roomRouter())
	r.Mount("/order", BH.orderRouter())
	return r
}
