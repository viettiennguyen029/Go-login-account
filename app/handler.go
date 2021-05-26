package app

import "github.com/gorilla/mux"

//IHandler ...
type IHandler interface {
	Attach(r *mux.Router)
}
