package route

import (
	"Zitank/models"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

func (BH BaseHandler) postRouter() http.Handler {
	r := chi.NewRouter()

	//get all post
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		posts, err := BH.postRepository.GetPosts()
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

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(BH.TokenAuth))
		r.Use(jwtauth.Authenticator(BH.TokenAuth))

		r.Route("/upload", func(r chi.Router) {
			r.Use(BH.AdminAuthenticate)

			r.Post("/", func(w http.ResponseWriter, r *http.Request) {
				var post models.Posts
				err := json.NewDecoder(r.Body).Decode(&post)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				err = BH.postRepository.CreatePost(&post)
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
				post, err := BH.postRepository.GetPost(postID)
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
			r.Group(func(r chi.Router) {
				r.Use(BH.AdminAuthenticate)
				r.Put("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					//check role
					_, claims, _ := jwtauth.FromContext(r.Context())
					userID := claims["userid"].(int)
					role := BH.userRepositor.RoleCheck(userID)
					if role != "admin" {
						http.Error(w, "Forbidden", http.StatusForbidden)
						return
					}

					var postValue models.Posts
					err := json.NewDecoder(r.Body).Decode(&postValue)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					BH.postRepository.UpdatePost(&postValue)
					w.Write([]byte("Successful update post"))
				}))

				r.Delete("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					//check role
					_, claims, _ := jwtauth.FromContext(r.Context())
					userID := claims["userid"].(int)
					role := BH.userRepositor.RoleCheck(userID)
					if role != "admin" {
						http.Error(w, "Forbidden", http.StatusForbidden)
						return
					}

					ctx := r.Context()
					postID, ok := ctx.Value("postID").(int)
					if !ok {
						http.Error(w, "Invalid post ID in context", http.StatusInternalServerError)
						return
					}
					BH.postRepository.DeletePost(postID)
					w.Write([]byte("Successful delete post"))
				}))
			})
		})
	})

	return r
}
