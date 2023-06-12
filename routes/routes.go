package routes

import (
	"net/http"

	"github.com/ojaoferreira/echo-server/controllers"
)

func HandlerRequests() {
	http.HandleFunc("/", controllers.GenericHandler)
}
