package main

import (
	_ "go-api/src/docs"
	"go-api/src/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
