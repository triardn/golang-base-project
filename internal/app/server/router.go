package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"github.com/triardn/golang-base-project/internal/app/handler"
)

// Router a chi mux
func Router(opt handler.HandlerOption) *chi.Mux {
	// cors
	corsOpt := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Content-Type", "X-Ktbs-Request-ID", "X-Ktbs-Api-Version", "X-Ktbs-Client-Version", "X-Ktbs-Platform-Name", "X-Ktbs-Client-Name", "Authorization", "X-Ktbs-Signature", "X-Ktbs-Time"},
	})

	healthCheckHandler := handler.NewHealtCheckHandler(opt)
	helloHandler := handler.NewHelloHandler(opt)

	r := chi.NewRouter()
	r.Use(corsOpt.Handler)

	// Setup your routing here
	r.Method(http.MethodGet, "/health/check", handler.HttpHandler{healthCheckHandler.Check})
	r.Method(http.MethodGet, "/health/readiness", handler.HttpHandler{healthCheckHandler.Readiness})
	r.Method(http.MethodGet, "/greet", handler.HttpHandler{helloHandler.Greet})

	return r
}

// TODO: func authRouter()
