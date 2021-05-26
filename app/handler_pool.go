package app

import (
	"github.com/gorilla/mux"
)

//IHandlerBuildable is inteface used to build handlers for server
type IHandlerBuildable interface {
	Build() *mux.Router
}

//HandlerPool containers IHandler and builds handlers for server
type HandlerPool struct {
	handlers []IHandler
}

//Push add new IHandler to pool
func (hp *HandlerPool) Push(h IHandler) {
	if hp.handlers == nil {
		hp.handlers = []IHandler{}
	}
	hp.handlers = append(hp.handlers, h)
}

//Build return a router for server
func (hp *HandlerPool) Build() *mux.Router {
	r := mux.NewRouter()
	for _, h := range hp.handlers {
		h.Attach(r)
	}
	return r
}
