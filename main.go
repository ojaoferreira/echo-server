package main

import (
	"fmt"
	"net/http"

	"github.com/ojaoferreira/echo-server/routes"
	"github.com/ojaoferreira/echo-server/utils"
)

var (
	Host string
	Port string
)

func init() {
	Host = utils.Getenv("HOST", "localhost")
	Port = utils.Getenv("PORT", "8080")
}

func main() {
	fmt.Printf("Iniciando o servidor. Access: http://%s:%s\n", Host, Port)

	routes.HandlerRequests()

	http.ListenAndServe(fmt.Sprintf(":%s", Port), nil)
}
