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

func musicRouter(rs *models.AppResource) http.Handler {
	r := chi.NewRouter()
	var musicController controller.MusicController
	var userController controller.UserController
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {

	})
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(rs.TokenAuth))
		r.Use(jwtauth.Authenticator(rs.TokenAuth))

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
				var music models.Musics
				err := json.NewDecoder(r.Body).Decode(&music)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				err = musicController.CreateMusic(rs.Store, music)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			})
		})

		r.Route("/{musicID}", func(r chi.Router) {
			r.Use(func(h http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					musicID, err := strconv.Atoi(chi.URLParam(r, "musicID"))
					if err != nil {
						http.Error(w, "Invalid music ID", http.StatusBadRequest)
						return
					}
					type contextKey string
					ctx := context.WithValue(r.Context(), contextKey("musicID"), musicID)
					h.ServeHTTP(w, r.WithContext(ctx))
				})
			})

			r.Post("/upload/info", func(w http.ResponseWriter, r *http.Request) {
				var musicinfo models.MusicInfo
				err := json.NewDecoder(r.Body).Decode(&musicinfo)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				err = musicController.CreateMusicInfo(rs.Store, musicinfo)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			})

			r.Put("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var musicValue models.Musics
				err := json.NewDecoder(r.Body).Decode(&musicValue)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				musicController.UpdateMusic(rs.Store, musicValue)
				w.Write([]byte("Successful update music infomation"))
			}))

			r.Delete("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				musicID, ok := ctx.Value("musicID").(int)
				if !ok {
					http.Error(w, "Invalid music ID in context", http.StatusInternalServerError)
					return
				}
				musicController.DeleteMusic(rs.Store, musicID)
				w.Write([]byte("Successful delete music infomation"))
			}))

			r.Route("/{musicInfoID}", func(r chi.Router) {
				r.Use(func(h http.Handler) http.Handler {
					return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						musicInfoID, err := strconv.Atoi(chi.URLParam(r, "musicInfoID"))
						if err != nil {
							http.Error(w, "Invalid music ID", http.StatusBadRequest)
							return
						}
						type contextKey string
						ctx := context.WithValue(r.Context(), contextKey("musicInfoID"), musicInfoID)
						h.ServeHTTP(w, r.WithContext(ctx))
					})
				})

				r.Put("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					var musicInfo models.MusicInfo
					err := json.NewDecoder(r.Body).Decode(&musicInfo)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					musicController.UpdateMusicInfo(rs.Store, musicInfo)
					w.Write([]byte("Successful update music info infomation"))
				}))

				r.Delete("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					ctx := r.Context()
					musicInfoID, ok := ctx.Value("musicInfoID").(int)
					if !ok {
						http.Error(w, "Invalid music info ID in context", http.StatusInternalServerError)
						return
					}
					musicController.DeleteMusicInfo(rs.Store, musicInfoID)
					w.Write([]byte("Successful delete music info infomation"))
				}))
			})

		})
	})
	r.Route("/{musicID}", func(r chi.Router) {
		r.Use(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				musicID, err := strconv.Atoi(chi.URLParam(r, "musicID"))
				if err != nil {
					http.Error(w, "Invalid music ID", http.StatusBadRequest)
					return
				}
				type contextKey string
				ctx := context.WithValue(r.Context(), contextKey("musicID"), musicID)
				h.ServeHTTP(w, r.WithContext(ctx))
			})
		})
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			musicID, ok := ctx.Value("musicID").(int)
			if !ok {
				http.Error(w, "Invalid music ID in context", http.StatusInternalServerError)
				return
			}
			musics, err := musicController.GetMusicsByID(rs.Store, musicID)
			if err != nil {
				http.Error(w, "Error fetching music", http.StatusInternalServerError)
				return
			}
			response, err := json.Marshal(musics)
			if err != nil {
				http.Error(w, "Error marshalling response", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		})
		r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			musicID, ok := ctx.Value("musicID").(int)
			if !ok {
				http.Error(w, "Invalid music ID in context", http.StatusInternalServerError)
				return
			}
			musics, err := musicController.GetMusicInfoByMusicID(rs.Store, musicID)
			if err != nil {
				http.Error(w, "Error fetching music info", http.StatusInternalServerError)
				return
			}
			response, err := json.Marshal(musics)
			if err != nil {
				http.Error(w, "Error marshalling response", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(response)
		})

	})
	r.Get("/all", func(w http.ResponseWriter, r *http.Request) {
		musics, err := musicController.GetMusics(rs.Store)
		if err != nil {
			http.Error(w, "Error fetching music", http.StatusInternalServerError)
			return
		}
		res, err := json.Marshal(musics)
		if err != nil {
			http.Error(w, "Error marshalling response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})
	return r
}
