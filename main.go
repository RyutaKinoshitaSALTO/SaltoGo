package main

import (
	_ "github.com/gin-gonic/gin"

	"SaltoGo/routes"
)

func main() {
	router := routes.NewRoutes()
	router.Run(":8080")
}
