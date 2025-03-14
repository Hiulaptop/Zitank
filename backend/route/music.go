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

func (BH BaseHandler) musicRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {

	})
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(BH.TokenAuth))
		r.Use(jwtauth.Authenticator(BH.TokenAuth))

		r.Route("/upload", func(r chi.Router) {
			r.Use(func(h http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// check role
					_, claims, _ := jwtauth.FromContext(r.Context())
					userID := claims["userid"].(int)
					role := BH.userRepositor.RoleCheck(userID)
					if role != "admin" {
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
				id, err := BH.musicRepository.CreateMusic(&music)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status": "success",
					"id":     id,
				})
			})
		})

		r.Route("/admin/{musicID}", func(r chi.Router) {
			r.Use(BH.AdminAuthenticate)

			r.Post("/upload/info", func(w http.ResponseWriter, r *http.Request) {
				var musicinfo models.MusicInfo
				err := json.NewDecoder(r.Body).Decode(&musicinfo)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				id, err := BH.musicRepository.CreateMusicInfo(&musicinfo)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status": "success",
					"id":     id,
				})
			})

			r.Put("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var musicValue models.Musics
				err := json.NewDecoder(r.Body).Decode(&musicValue)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				BH.musicRepository.UpdateMusic(&musicValue)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status": "success",
				})
			}))

			r.Delete("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				musicID, ok := ctx.Value("musicID").(int)
				if !ok {
					http.Error(w, "Invalid music ID in context", http.StatusInternalServerError)
					return
				}
				BH.musicRepository.DeleteMusic(musicID)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"status": "success",
				})
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
					BH.musicRepository.UpdateMusicInfo(&musicInfo)
					json.NewEncoder(w).Encode(map[string]interface{}{
						"status": "success",
					})
				}))

				r.Delete("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					ctx := r.Context()
					musicInfoID, ok := ctx.Value("musicInfoID").(int)
					if !ok {
						http.Error(w, "Invalid music info ID in context", http.StatusInternalServerError)
						return
					}
					BH.musicRepository.DeleteMusicInfo(musicInfoID)
					json.NewEncoder(w).Encode(map[string]interface{}{
						"status": "success",
					})
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
			musics, err := BH.musicRepository.GetMusicsByID(musicID)
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
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status": "success",
				"music":  response,
			})
		})
		r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			musicID, ok := ctx.Value("musicID").(int)
			if !ok {
				http.Error(w, "Invalid music ID in context", http.StatusInternalServerError)
				return
			}
			musics, err := BH.musicRepository.GetMusicInfoByMusicID(musicID)
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
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":     "success",
				"music_info": response,
			})
		})

	})
	r.Get("/all", func(w http.ResponseWriter, r *http.Request) {
		musics, err := BH.musicRepository.GetMusics()
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
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"musics": res,
		})
	})
	return r
}
