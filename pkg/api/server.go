package api

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	pathPrefix = "/polmon/api"
)

type HttpServer struct {
	conf    Configure
	router  *mux.Router
	srv     *http.Server
	handler Handler
}

func NewServer(conf Configure, handler Handler) *HttpServer {
	return &HttpServer{
		conf:    conf,
		router:  mux.NewRouter(),
		handler: handler,
	}
}

func (h *HttpServer) Start() {
	h.router.StrictSlash(true)
	s := h.router.PathPrefix(pathPrefix).Subrouter()
	h.initApiV1(s)

	h.srv = &http.Server{
		Addr: h.conf.Addr,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: h.conf.WriteTimeout,
		ReadTimeout:  h.conf.ReadTimeout,
		IdleTimeout:  h.conf.IdleTimeout,
		Handler:      h.router, // Pass our instance of gorilla/mux in.
	}

	log.Printf("<INF> http listening %s\n", h.conf.Addr)
	go func() {
		if err := h.srv.ListenAndServe(); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()
}

func (h *HttpServer) initApiV1(r *mux.Router) {
	s := r.PathPrefix("/v1").Subrouter()
	s.HandleFunc("/fee", h.handler.Fee).Methods("GET")
}

func (h *HttpServer) Stop() {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), h.conf.ShutdownTimeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	h.srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
}
