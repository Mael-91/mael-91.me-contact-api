package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type Handlers struct {
	Template *template.Template
}

func serve() {
	handlers := &Handlers{Template: initTemplate()}
	srv := http.Server{
		Addr: ":2000",
		Handler: handlers.GetRoutes(),
	}
	fmt.Printf("Starting HTTP server on port %s\n", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (h *Handlers) GetRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/contact", h.contactFormHandler).
		Methods(http.MethodPost, http.MethodOptions)
	r.Use(CORSMethodMiddleware("/contact"))
	return r
}

func CORSMethodMiddleware(route string) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost && r.RequestURI == route {
				for key, value := range map[string]string{
					"Access-Control-Allow-Origin":  "https://mael-91.me",
					"Access-Control-Allow-Methods": "OPTIONS, POST",
					"Access-Control-Allow-Headers": "Content-Type, Accept",
				} {
					w.Header().Set(key, value)
				}
			}
			handler.ServeHTTP(w, r)
		})
	}
}
