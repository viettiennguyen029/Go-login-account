package app

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (

	//SECRET_JWT SHOULD MOVE TO CONFIG
	SECRET_JWT = "$2a$14$hNjUMdc8oRvfd8DAh6Z8Ye8JM4kBpCz597YZajnEnLGUPf0ia1TR2"
)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}
