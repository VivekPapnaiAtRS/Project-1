package server

import (
	"github.com/Project/handler"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server struct {
	chi.Router
}

func SetupRoutes() *Server {
	router := chi.NewRouter()
	router.Route("/api", func(api chi.Router) {
		api.Get("/", func(responseWriter http.ResponseWriter, request *http.Request) {
			_, err := responseWriter.Write([]byte("So you are on vivek's first deployed server.\n\n\t you must be feeling lucky"))
			if err != nil {
				return
			}
		})
		api.Post("/text", handler.CountFrequency)
	})
	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
