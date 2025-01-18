package route

import (
	"Zitank/controller"
	"Zitank/models"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func postRouter(rs *models.AppResource) http.Handler {
	r := chi.NewRouter()
	var userController controller.UserController
	var postController controller.PostController

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {

	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(rs.TokenAuth))
		r.Use(jwtauth.Authenticator(rs.TokenAuth))

		r.Get("/all", func(w http.ResponseWriter, r *http.Request) {
			posts, err := postController.GetPosts(rs.Store)
			if err != nil {
				http.Error(w, "Error fetching post", http.StatusInternalServerError)
				return
			}
			response, err := json.Marshal(posts)
			if err != nil {
				http.Error(w, "Error marshalling response", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		})

		// get by user name

		r.Route("/upload", func(r chi.Router) {
			r.Use(func(h http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// check role
					_, claims, _ := jwtauth.FromContext(r.Context())
					userID := claims["user_id"].(int)
					ok := userController.RoleCheck(rs.Store, userID)
					if !ok {
						http.Error(w, "Forbidden", http.StatusForbidden)
						return
					}
					h.ServeHTTP(w, r)
				})
			})

			r.Post("/", func(w http.ResponseWriter, r *http.Request) {
				var post models.Posts
				err := json.NewDecoder(r.Body).Decode(&post)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				err = postController.CreatePost(rs.Store, post)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			})
		})

		r.Route("/{postID}", func(r chi.Router) {
			r.Use(func(h http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					postID, err := strconv.Atoi(chi.URLParam(r, "postID"))
					if err != nil {
						http.Error(w, "Invalid post ID", http.StatusBadRequest)
						return
					}
					type contextKey string
					ctx := context.WithValue(r.Context(), contextKey("postID"), postID)
					h.ServeHTTP(w, r.WithContext(ctx))
				})
			})

			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				postID, ok := ctx.Value("postID").(int)
				if !ok {
					http.Error(w, "Invalid post ID in context", http.StatusInternalServerError)
					return
				}
				post, err := postController.GetPost(rs.Store, postID)
				if err != nil {
					http.Error(w, "Error fetching post", http.StatusInternalServerError)
					return
				}
				response, err := json.Marshal(post)
				if err != nil {
					http.Error(w, "Error marshalling response", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(response)
			})
		})
	})

	return r
}
