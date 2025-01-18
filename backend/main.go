package main

import (
	"Zitank/models"
	"Zitank/route"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/justinas/nosurf"
	"github.com/unrolled/secure"
)

func main() {
	// DOTENV
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// SQLX
	var db models.PostgresStore
	err = db.Connect(os.Getenv("DB_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// JWT
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)

	// CHI ROUTER
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:          []string{"example\\.com", ".*\\.example\\.com"},
		AllowedHostsAreRegex:  true,
		HostsProxyHeaders:     []string{"X-Forwarded-Host"},
		SSLRedirect:           true,
		SSLHost:               "ssl.example.com",
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:            31536000,
		STSIncludeSubdomains:  true,
		STSPreload:            true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "script-src $NONCE",
		IsDevelopment:         true,
	})
	r.Use(secureMiddleware.Handler)
	r.Use(nosurf.NewPure)

	rs := models.NewAppResource(&db, tokenAuth)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	r.Mount("/api", route.ApiRouter(rs))
}
