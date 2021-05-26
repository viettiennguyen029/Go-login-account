package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

//UserHandler provides function handlers of user
type UserHandler struct {
}

//Attach login function to router
func (uh *UserHandler) Attach(r *mux.Router) {
	r.HandleFunc("/users/login", uh.login).Methods("POST")
	r.HandleFunc("/users/signup", uh.signup).Methods("POST")
}

//Signup ...
func (uh *UserHandler) signup(w http.ResponseWriter, r *http.Request) {

	// Example for sending a error response
	// NewJSendJSONBuilder().
	// 		Code(http.StatusBadRequest).
	// 		Message(err.Error()).
	// 		Build().
	// 		Send(w)

	// Example for sending success reponse
	// NewJSendJSONBuilder().
	// 	Code(http.StatusOK).
	// 	Build().
	// 	Send(w)
}
func (uh *UserHandler) login(w http.ResponseWriter, r *http.Request) {

}
