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

func (BH BaseHandler) orderRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(jwtauth.Verifier(BH.TokenAuth))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		type GetBody struct {
			Status string           `json:"status"`
			Orders []*models.Orders `json:"orders"`
		}
		_, claims, _ := jwtauth.FromContext(r.Context())
		userID, ok := claims["userid"]
		if !ok || BH.userRepositor.RoleCheck(int(userID.(float64))) != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("401 - Unauthorized"))
			return
		}
		orders, err := BH.orderRepository.GetOrders()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		var body GetBody
		body.Status = "success"
		body.Orders = orders
		res, err := json.Marshal(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(res)
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Authenticator(BH.TokenAuth))

		r.Route("/{orderID}", func(r chi.Router) {
			r.Use(func(h http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					orderID, err := strconv.Atoi(chi.URLParam(r, "orderID"))
					if err != nil {
						http.Error(w, "Invalid order ID", http.StatusBadRequest)
						return
					}
					type contextKey string
					ctx := context.WithValue(r.Context(), contextKey("orderID"), orderID)
					h.ServeHTTP(w, r.WithContext(ctx))
				})
			})

			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				ctx := r.Context()
				orderID, ok := ctx.Value("orderID").(int)
				if !ok {
					http.Error(w, "Invalid order ID in context", http.StatusInternalServerError)
					return
				}
				order, err := BH.orderRepository.GetOrder(orderID)
				if err != nil {
					http.Error(w, "Error fetching order", http.StatusInternalServerError)
					return
				}
				response, err := json.Marshal(order)
				if err != nil {
					http.Error(w, "Error marshalling response", http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(map[string]any{
					"status": "success",
					"order":  response,
				})
			})
			r.Group(func(r chi.Router) {
				r.Use(BH.AdminAuthenticate)

				r.Put("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					var orderValue models.Orders
					err := json.NewDecoder(r.Body).Decode(&orderValue)
					if err != nil {
						http.Error(w, err.Error(), http.StatusBadRequest)
						return
					}
					BH.orderRepository.UpdateOrder(&orderValue)
					json.NewEncoder(w).Encode(map[string]interface{}{
						"status": "success",
					})
				}))

				r.Delete("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					ctx := r.Context()
					orderID, ok := ctx.Value("orderID").(int)
					if !ok {
						http.Error(w, "Invalid order ID in context", http.StatusInternalServerError)
						return
					}
					BH.orderRepository.DeleteOrder(orderID)
					json.NewEncoder(w).Encode(map[string]interface{}{
						"status": "success",
					})
				}))
			})
		})
	})

	return r
}
