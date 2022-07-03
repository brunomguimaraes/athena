package api

import (
	"athena/api/controllers"
)

var server = controllers.Server{}

func Start() {
	server.Initialize()

	server.Run(":8080")
}
