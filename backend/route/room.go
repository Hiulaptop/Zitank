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

func roomRouter(rs *models.AppResource) http.Handler {
	r := chi.NewRouter()
	var userController controller.UserController
	var roomController controller.RoomController

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		rooms, err := roomController.GetRooms(rs.Store)
		if err != nil {
			http.Error(w, "Error fetching room", http.StatusInternalServerError)
			return
		}
		response, err := json.Marshal(rooms)
		if err != nil {
			http.Error(w, "Eror marshalling response", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(rs.TokenAuth))
		r.Use(jwtauth.Authenticator(rs.TokenAuth))

		r.Route("/upload", func(r chi.Router){
			r.Use(func(h http.Handler) http.Handler{
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					//check role
					_, claims, _ := jwtauth.FromContext(r.Context())
					userID := claims["user_id"].(int)
					ok := userController.RoleCheck(rs.Store, userID)
					if !ok {
						http.Error(w, "Forbiden", http.StatusForbidden)
						return
					}

					h.ServeHTTP(w, r)
				})
			})

			r.Post("/", func(w http.ResponseWriter, r *http.Request) {
				var room models.Rooms
				err := json.NewDecoder(r.Body).Decode(&room)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				err = roomController.CreateRoom(rs.Store, room)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			})
		})

		r.Route("/{roomID}", func(r chi.Router){
			r.Use(func(h http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					roomID, err := strconv.Atoi(chi.URLParam(r, "roomID"))
					if err != nil {
						http.Error(w, "Invalid room ID", http.StatusBadRequest)
						return
					}
					type contextKey string
					ctx := context.WithValue(r.Context(), contextKey("roomID"), roomID)
					h.ServeHTTP(w, r.WithContext(ctx))
				})
			})

			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				roomID, ok := ctx.Value("roomID").(int)
				if !ok {
					http.Error(w, "Invalid room ID in context", http.StatusInternalServerError)
					return
				}
				room, err := roomController.GetRoom(rs.Store, roomID)
				if err != nil {
					http.Error(w, "Error fetching room", http.StatusInternalServerError)
					return
				}
				response, err := json.Marshal(room)
				if err != nil {
					http.Error(w, "Error marshalling response", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write(response)
			})

			r.Put("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				//check role
				_, claims, _ := jwtauth.FromContext(r.Context())
				userID := claims["user_id"].(int)
				ok := userController.RoleCheck(rs.Store, userID)
				if !ok {
					http.Error(w, "Forbidden", http.StatusForbidden)
					return
				}

				var roomValue models.Rooms
				err := json.NewDecoder(r.Body).Decode(&roomValue)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				roomController.UpdateRoom(rs.Store, roomValue)
				w.Write([]byte("Successful update room"))
			}))

			r.Delete("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				//check role
				_, claims, _ := jwtauth.FromContext(r.Context())
				userID := claims["user_id"].(int)
				ok := userController.RoleCheck(rs.Store, userID)
				if !ok {
					http.Error(w, "Forbidden", http.StatusForbidden)
					return
				}

				ctx := r.Context()
				roomID, ok := ctx.Value("roomID").(int)
				if !ok {
					http.Error(w, "Invalid room ID in context", http.StatusInternalServerError)
					return
				}
				roomController.DeleteRoom(rs.Store, roomID)
				w.Write([]byte("Successful delete room"))
			}))
		})
	})

	return r
}
