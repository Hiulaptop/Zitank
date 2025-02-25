package route

import (
	// "Zitank/models"

	"Zitank/models"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/pgtype"
)

type GetFreeTimeBody struct {
	FromTo string `json:"fromto"`
}

type OrderTime struct {
	From   time.Time
	To     time.Time
	Status string
}

type FreeTime struct {
	FromTo []OrderTime `json:"fromto"`
}

type OrderBody struct {
	FromTo string `json:"fromto"`
	State  string `json:"state"`
	Note   string `json:"note"`
	UserID uint   `json:"userid"`
	RoomID uint   `json:"roomid"`
}

func fromOrderBody(OB OrderBody) (models.Orders, error) {
	var FT pgtype.Tsrange
	err := FT.Scan(OB.FromTo)
	return models.Orders{
		FromTo: FT,
		State:  OB.State,
		Note:   OB.Note,
		// TotalPrice: float64(OB.TotalPrice),
		UserID: OB.UserID,
		RoomID: OB.RoomID,
	}, err
}

// rs *models.AppResource
func (BH BaseHandler) roomRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		rooms, err := BH.roomRepository.GetRooms()
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
		r.Use(jwtauth.Verifier(BH.TokenAuth))
		r.Use(jwtauth.Authenticator(BH.TokenAuth))

		r.Route("/upload", func(r chi.Router) {
			r.Use(func(h http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					ctx := r.Context()
					//check role
					_, claims, _ := jwtauth.FromContext(ctx)
					userID := int(math.Round(claims["userid"].(float64)))
					role := BH.userRepositor.RoleCheck(userID)
					if role != "admin" {
						http.Error(w, "Forbidden", http.StatusForbidden)
						return
					}
					ctx = context.WithValue(ctx, "userid", userID)
					h.ServeHTTP(w, r.WithContext(ctx))
				})
			})

			r.Post("/", func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				userID, ok := ctx.Value("userid").(int)
				if !ok {
					http.Error(w, "Unknown User.", http.StatusBadRequest)
					return
				}
				var room models.RoomObject
				err := json.NewDecoder(r.Body).Decode(&room)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				err = BH.roomRepository.CreateRoom(&room, userID)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			})
		})

		r.Route("/{roomID}", func(r chi.Router) {
			r.Use(func(h http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					roomID, err := strconv.Atoi(chi.URLParam(r, "roomID"))
					if err != nil {
						http.Error(w, "Invalid room ID", http.StatusBadRequest)
						// w.Write([]byte("Hello, World!"))
						return
					}
					ctx := context.WithValue(r.Context(), "roomID", roomID)
					_, claims, _ := jwtauth.FromContext(ctx)
					userID := int(math.Round(claims["userid"].(float64)))
					ctx = context.WithValue(ctx, "userid", userID)
					h.ServeHTTP(w, r.WithContext(ctx))
				})
			})

			//need to edit
			r.Post("/order", func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				roomID, ok := ctx.Value("roomID").(int)
				if !ok {
					http.Error(w, "Invalid room ID in context", http.StatusInternalServerError)
					return
				}
				userID, ok := ctx.Value("userid").(int)
				if !ok {
					http.Error(w, "Invalid room ID in context", http.StatusInternalServerError)
					return
				}
				var OB OrderBody
				err := json.NewDecoder(r.Body).Decode(&OB)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				OB.UserID = uint(userID)
				OB.RoomID = uint(roomID)
				order, err := fromOrderBody(OB)
				//Fix me
				//Fix me
				//Fix me
				//Fix me
				//Fix me
				//Fix me
				//Fix me
				order.TotalPrice = 12.3
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				err = BH.orderRepository.CreateOrder(&order)
				if err != nil {
					fmt.Println(err)
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			})

			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				// w.Write([]byte("Hello, World!"))
				ctx := r.Context()
				roomID, ok := ctx.Value("roomID").(int)
				if !ok {
					http.Error(w, "Invalid room ID in context", http.StatusInternalServerError)
					return
				}
				room, err := BH.roomRepository.GetRoom(roomID)
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
				userID := claims["userid"].(int)
				role := BH.userRepositor.RoleCheck(userID)
				if role != "admin" {
					http.Error(w, "Forbidden", http.StatusForbidden)
					return
				}

				var roomValue models.Rooms
				err := json.NewDecoder(r.Body).Decode(&roomValue)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				BH.roomRepository.UpdateRoom(&roomValue)
				w.Write([]byte("Successful update room"))
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
				roomID, ok := ctx.Value("roomID").(int)
				if !ok {
					http.Error(w, "Invalid room ID in context", http.StatusInternalServerError)
					return
				}
				BH.roomRepository.DeleteRoom(roomID)
				w.Write([]byte("Successful delete room"))
			}))

			r.Post("/getfreetime", func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				roomID, ok := ctx.Value("roomID").(int)
				if !ok {
					http.Error(w, "Invalid room ID in context", http.StatusInternalServerError)
					return
				}
				var fromto GetFreeTimeBody
				var ts pgtype.Tsrange
				err := json.NewDecoder(r.Body).Decode(&fromto)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				err = ts.Scan(fromto.FromTo)
				if err != nil {
					fmt.Println(err)
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				orders, err := BH.orderRepository.GetOrdersInRange(roomID, ts)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				var fr FreeTime
				for _, ele := range orders {
					var OT OrderTime
					OT.To = ele.FromTo.Upper.Time
					OT.From = ele.FromTo.Lower.Time
					OT.Status = "In process"
					fr.FromTo = append(fr.FromTo, OT)
				}
				json.NewEncoder(w).Encode(fr)
			})
		})
	})

	return r
}
