package main

import (
	"github.com/ameykpatil/wanderlife/handlers"
	"github.com/gin-gonic/gin"
	"gopkg.in/danilopolani/gocialite.v0"
)

var gocial = gocialite.NewDispatcher()

func main() {
	router := gin.Default()

	router.GET("/", handlers.LoginIndex)
	router.GET("/auth/:provider", handlers.LoginRedirect)
	router.GET("/auth/:provider/callback", handlers.LoginCallback)

	router.Run("127.0.0.1:9090")
}
