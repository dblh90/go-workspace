package main

import (
	"github.com/gin-gonic/gin"

	"go-fullstack/internals"
)

func main() {
	router := gin.Default()

	app := internals.Config{
		Router: router,
	}

	app.Routes()
	router.Run(":9090")
}
