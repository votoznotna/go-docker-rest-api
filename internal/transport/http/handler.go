package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service CommentService
	Server  *http.Server
}

// Response objecgi
type Response struct {
	Message string `json:"message"`
}

func NewHandler(service CommentService) *Handler {
	log.Info("setting up our handler")
	h := &Handler{
		Service: service,
	}

	h.Router = mux.NewRouter()
	// Sets up our middleware functions
	h.Router.Use(JSONMiddleware)
	// we also want to log every incoming request
	h.Router.Use(LoggingMiddleware)
	// We want to timeout all requests that take longer than 15 seconds
	h.Router.Use(TimeoutMiddleware)
	// set up the routes
	h.mapRoutes()

	h.Server = &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		// WriteTimeout: time.Second * 15,
		// ReadTimeout:  time.Second * 15,
		// IdleTimeout:  time.Second * 60,
		Handler: h.Router,
	}
	// // return our wonderful handler
	return h
}

// mapRoutes - sets up all the routes for our application
func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	h.Router.HandleFunc("/alive", h.AliveCheck).Methods("GET")
	// h.Router.HandleFunc("/ready", h.ReadyCheck).Methods("GET")
	h.Router.HandleFunc("/api/v1/comment", JWTAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/v1/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/v1/comment/{id}", JWTAuth(h.UpdateComment)).Methods("PUT")
	h.Router.HandleFunc("/api/v1/comment/{id}", JWTAuth(h.DeleteComment)).Methods("DELETE")

}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// // Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shutting down gracefully")

	return nil
}
