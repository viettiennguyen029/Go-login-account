package app

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

//IServiceServer interface for servers
type IServiceServer interface {
	Start()
}

//ServiceServer is the server serve api for the block explorer of tron based network
type ServiceServer struct {
	server         *http.Server
	handlerBuilder IHandlerBuildable
}

func (s *ServiceServer) buildRouter() {
	r := s.handlerBuilder.Build()
	//sse
	//CORS
	s.server.Handler = handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(r)
	r.Use(setDefaultHeadersForRes)
}
func setDefaultHeadersForRes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

//Start run server
func (s *ServiceServer) Start() {
	s.buildRouter()
	log.Print("Started")
	s.server.ListenAndServe()
}

//NewServiceServer return new server
func NewServiceServer(addr string, handlerBuilder IHandlerBuildable) IServiceServer {
	return &ServiceServer{
		server: &http.Server{
			Addr: addr,
		},
		handlerBuilder: handlerBuilder,
	}
}
