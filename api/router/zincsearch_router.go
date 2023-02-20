package router

import (
	chi "github.com/go-chi/chi"
	middleware "github.com/go-chi/chi/v5/middleware"
	cors "github.com/go-chi/cors"
	jwtauth "github.com/go-chi/jwtauth/v5"

	zincSearch "falconEmailBackend/api/handler/zincsearch/search"
)

// Vars
var (
	app       = chi.NewRouter()
	tokenAuth *jwtauth.JWTAuth
)

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
	_, _, _ = tokenAuth.Encode(map[string]interface{}{"user": 123})
}

// InitializeZincSearchRouter function that return a server using chi library
func InitializeZincSearchRouter() *chi.Mux {
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Use(middleware.RedirectSlashes)
	app.Use(middleware.StripSlashes)
	app.Use(middleware.AllowContentType("application/json", "text/xml"))
	app.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "HEAD", "OPTION"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// routes protected
	app.Group(func(app chi.Router) {
		app.Use(jwtauth.Verifier(tokenAuth))
		app.Use(jwtauth.Authenticator)
		app.Mount("/zinsearch/search", zincSearch.ZincSearchRoute())
	})

	// public routes
	app.Group(func(r chi.Router) {
		// app.Mount("/zinsearch/search", zincSearch.ZincSearchRoute())
	})

	return app
}
