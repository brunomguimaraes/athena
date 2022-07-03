package main

import (
	"athena/api"

	_ "github.com/lib/pq"
)

func main() {
	api.Start()
}
